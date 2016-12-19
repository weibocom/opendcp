#!/usr/bin/env python
#
#    Copyright (C) 2016 Weibo Inc.
#
#    This file is part of Opendcp.
#
#    Opendcp is free software: you can redistribute it and/or modify
#    it under the terms of the GNU General Public License as published by
#    the Free Software Foundation; version 2 of the License.
#
#    Opendcp is distributed in the hope that it will be useful,
#    but WITHOUT ANY WARRANTY; without even the implied warranty of
#    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
#    GNU General Public License for more details.
#
#    You should have received a copy of the GNU General Public License
#    along with Opendcp.  if not, write to the Free Software
#    Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA 02110-1301  USA
#
#
# -*- coding: utf-8 -*-

# Author: WhiteBlue
# Time  : 2016/07/20


from __future__ import (absolute_import, division, print_function)

import time
import json
import uuid
import datetime
from ansible.utils.unicode import to_bytes
from ansible.plugins.callback import CallbackBase
from octans.logger import LogManager
from octans import Service

Logger = LogManager.get_logger("SyncCallbackModule")
__metaclass__ = type

FAILED = -1
UNREACHABLE = -2
ASYNC_FAILED = -3
OK = 0


def GetJSONProperty(req_json, key):
    if req_json.has_key(key):
        value = req_json[key]
    return value


class SyncCallbackModule(CallbackBase):
    """
    logs playbook results, per host
    """
    CALLBACK_VERSION = 2.0
    CALLBACK_TYPE = 'notification'
    CALLBACK_NAME = 'log_plays'
    CALLBACK_NEEDS_WHITELIST = True

    TIME_FORMAT = "%b %d %Y %H:%M:%S"
    MSG_FORMAT = "%(now)s - %(category)s - %(data)s\n\n"

    def __init__(self, step_callback, global_id, source, tag_hosts, debug=False):
        super(SyncCallbackModule, self).__init__()
        self._debug = debug
        self.global_id = global_id
        self.source = source
        self.task_uuid = None
        self.task_status = None
        self._step_callback = step_callback
        self.tag_hosts = tag_hosts

    def log(self, host, code, data):
     
        '''
        if type(data) == dict:
            if '_ansible_verbose_override' in data:
                data = 'omitted'
            else:
                data = data.copy()
                invocation = data.pop('invocation', None)
                data = json.dumps(data)
                if invocation is not None:
                    data = json.dumps(invocation) + " => %s " % data
        '''

        newdata={};
        if data.has_key('cmd'):
            newdata['cmd']=data['cmd']
        if data.has_key('start'):
            newdata['start']=data['start']
        if data.has_key('end'):
            newdata['end']=data['end']
        if data.has_key('delta'):
            newdata['delta']=data['delta']
        if data.has_key('stderr'):
            newdata['stderr']=data['stderr']
        if data.has_key('stdout'):
            newdata['stdout']=data['stdout']
        if data.has_key('stdout_lines'):
            newdata['stdout_lines']=data['stdout_lines']
        
        self._step_callback(ip=host, code=code, data=newdata)
        
        if self._debug:
            now = time.strftime(self.TIME_FORMAT, time.localtime())
            msg = to_bytes(self.MSG_FORMAT % dict(now=now, category=code, data=data))
            print(msg)

    def playbook_on_task_start(self, name, is_conditional):
        self.task_uuid = uuid.uuid1().hex
        create_time = datetime.datetime.now()
        create_time = create_time.strftime('%Y-%m-%dT%H:%M:%SZ%z')
        for tag in self.tag_hosts:
            Service.add_log(global_id=self.global_id, source=self.source, task_uuid=self.task_uuid, task_status="", create_time=create_time, end_time="",data="",host=tag )

    def runner_on_failed(self, host, res, ignore_errors=False):
        Logger.debug("fail -----------------------------------")
        a = json.dumps(res)  
        Logger.debug(a)
        end_time = datetime.datetime.now()                                 
        end_time = end_time.strftime('%Y-%m-%dT%H:%M:%SZ%z')
        self.log(host, FAILED, res)
        Service.update_log(host=host, global_id=self.global_id, task_uuid=self.task_uuid, task_status="failed",end_time=end_time, data=a) 

    def runner_on_ok(self, host, res):
        
        Logger.debug("success **********************************")
        a = json.dumps(res)  
        Logger.debug(a)
        end_time = datetime.datetime.now()
        end_time = end_time.strftime('%Y-%m-%dT%H:%M:%SZ%z')        
        self.log(host, OK, res)
        Service.update_log(host=host, global_id=self.global_id, task_uuid=self.task_uuid, task_status="ok", end_time=end_time,data=a )

    def runner_on_unreachable(self, host, res):
        Logger.debug("unreachable -----------------------------------")
        a = json.dumps(res)  
        Logger.debug(a)
        end_time = datetime.datetime.now()
        end_time = end_time.strftime('%Y-%m-%dT%H:%M:%SZ%z')        
        self.log(host, UNREACHABLE, res)
        Service.update_log(host=host, global_id=self.global_id, task_uuid=self.task_uuid, task_status="unreachable", end_time=end_time,data=a )


    def runner_on_async_failed(self, host, res, jid):
        Logger.debug("async fail -----------------------------------")
        end_time = datetime.datetime.now()
        end_time = end_time.strftime('%Y-%m-%dT%H:%M:%SZ%z')        
        a = json.dumps(res)  
        Logger.debug(a)

        self.log(host, ASYNC_FAILED, res)
        Service.update_log(host=host, global_id=self.global_id, task_uuid=self.task_uuid, task_status="async_failed", end_time=end_time,data=a )

