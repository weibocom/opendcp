# Ansible is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# Ansible is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with Ansible.  If not, see <http://www.gnu.org/licenses/>.

from __future__ import (absolute_import, division, print_function)

__metaclass__ = type

import uuid
import time
import json
from ansible.plugins.callback import CallbackBase
from octans.logger import LogManager
from octans import Service
from ansible.errors import AnsibleError

Logger = LogManager.get_logger('NewCallbackModule')

class NewCallbackModule(CallbackBase):

    START = "start"
    FAILED = "failed"
    SUCCESS = "ok"
    UNREACHABLE = "unreachable"
    SKIP = "skip"

    TIME_FORMAT = '%Y-%m-%d %H:%M:%S'
    MSG_FORMAT = '%(now)s - %(category)s - %(data)s\n\n'

    def __init__(self, host, step_callback, global_id, source, task_id):
        super(NewCallbackModule, self).__init__()
        self.uid = uuid.uuid1().hex
        self.host = host
        self._display.verbosity = 3
        self.max_try = 3
        self._step_callback = step_callback
        self.global_id = global_id
        self.source = source
        self.task_id = task_id

    def v2_runner_on_ok(self, result):
        task = result._task
        self._display.display(
            ('end - OK - %s: %s - %s' % (task.get_name(), getattr(task, '_uuid'), self.host)),
            'green')
        Logger.info('>>>>>OK {}'.format(result._result))
        task_info = 'Finish run task successfully'
        self.update_log(task_status=self.SUCCESS, detail=task_info)

    def v2_runner_on_skipped(self, result):
        task = result._task
        self._display.display(
            ('end - Skip - %s: %s - %s' % (task.get_name(), getattr(task, '_uuid'), self.host)),
            'cyan')
        task_info = 'runner on skipped'
        Logger.info('>>>>>Skip {}'.format(task_info))
        self.update_log(task_status=self.SKIP, detail=result._result)

    def v2_runner_on_failed(self, result, ignore_errors=False):
        task = result._task
        self._display.display(
            ('end - Failed - %s: %s - %s' % (task.get_name(), getattr(task, '_uuid'), self.host)),
            'red')
        task_info = 'runner on failed'
        Logger.info('>>>>>Failed {}'.format(task_info))

        self.update_log(task_status=self.FAILED, detail=self._dump_results(result._result, indent=4))

    def v2_runner_on_unreachable(self, result):
        task = result._task
        self._display.display(
            ('end - Unreachable - %s: %s - %s' % (
                task.get_name(), getattr(task, '_uuid'), self.host)),
            'bright red')
        Logger.info('>>>>>Unreachable {}'.format(result._result))
        task_info = 'Finish run task with error unreachable'
        self.update_log(task_status=self.UNREACHABLE, detail=task_info)

    def v2_playbook_on_task_start(self, task, is_conditional):
        action_name = task.get_name().strip()
        setattr(task, '_uuid', self.uid)
        self._display.display(('start - %s: %s - %s' % (action_name, self.uid, self.host)),
                              color='dark gray')
        if task.action == 'include':
            return
        task_vars = task.get_vars()
        task_module = task.action
        task_params = task.get_include_params()
        task_path = task.get_path()
        task_dict = dict(task_vars=task_vars, task_module=task_module, task_params=task_params,
                         task_path=task_path)
        task_info = 'Begin run task......'
        Logger.info('>>>>>{}'.format(task_info))
        Logger.info('>>>>>Start, the params is: {}', json.dumps(task_dict))
        self.add_log(task_status=self.START, detail=task_info)

    def add_log(self, task_status, detail):
        begin_time = time.strftime(self.TIME_FORMAT, time.localtime())
        for i in range(1, self.max_try + 1):
            try:
                Service.add_log(global_id=self.global_id, source=self.source, task_uuid=self.uid, task_status=task_status,
                                create_time=begin_time, end_time="",data= detail, host=self.host, task_id = self.task_id)
                return
            except Exception as err:
                Logger.error('Add log err for %d times: {}'.format(str(err)) % i)
        raise AnsibleError("Add log DB error retry for %d times" % self.max_try)

    def update_log(self, task_status, detail):
        finish_time = time.strftime(self.TIME_FORMAT, time.localtime())
        for i in range(1, self.max_try + 1):
            try:
                Service.update_log(host=self.host, global_id=self.global_id, task_uuid=self.uid, task_status=task_status,
                                   end_time=finish_time, data=detail)
                return
            except Exception as err:
                Logger.error("Update log err for %d times:{}".format(str(err)) % i)
        raise AnsibleError("Update log DB error retry for %d times" % self.max_try)