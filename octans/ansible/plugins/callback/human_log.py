# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with this program.  If not, see <http://www.gnu.org/licenses/>.

from ansible.plugins.callback import CallbackBase

FIELDS = ['cmd', 'command', 'start', 'end', 'delta', 'msg', 'stdout', 'stderr']

def human_log(res):

    if type(res) == type(dict()):
        for field in FIELDS:
            if field in res.keys():
                # use default encoding, check out sys.setdefaultencoding
                print u'\n{0}:\n{1}'.format(field, res[field])
                # or use specific encoding, e.g. utf-8
                #print '\n{0}:\n{1}'.format(field, res[field].encode('utf-8'))

class CallbackModule(object):

    def on_any(self, *args, **kwargs):
        pass

    def runner_on_failed(self, host, res, ignore_errors=False):
        human_log(res)

    def runner_on_ok(self, host, res):
        human_log(res)

    def runner_on_error(self, host, msg):
        pass

    def runner_on_skipped(self, host, item=None):
        pass

    def runner_on_unreachable(self, host, res):
        human_log(res)

    def runner_on_no_hosts(self):
        pass

    def runner_on_async_poll(self, host, res, jid, clock):
        human_log(res)

    def runner_on_async_ok(self, host, res, jid):
        human_log(res)

    def runner_on_async_failed(self, host, res, jid):
        human_log(res)

    def playbook_on_start(self):
        pass

    def playbook_on_notify(self, host, handler):
        pass

    def playbook_on_no_hosts_matched(self):
        pass

    def playbook_on_no_hosts_remaining(self):
        pass

    def playbook_on_task_start(self, name, is_conditional):
        pass

    def playbook_on_vars_prompt(self, varname, private=True, prompt=None, encrypt=None, confirm=False, salt_size=None, salt=None, default=None):
        pass

    def playbook_on_setup(self):
        pass

    def playbook_on_import_for_host(self, host, imported_file):
        pass

    def playbook_on_not_import_for_host(self, host, missing_file):
        pass

    def playbook_on_play_start(self, pattern):
        pass

    def playbook_on_stats(self, stats):
        pass