#!/usr/bin/env python
#
#  Copyright 2009-2016 Weibo, Inc.
#
#    Licensed under the Apache License, Version 2.0 (the "License");
#    you may not use this file except in compliance with the License.
#    You may obtain a copy of the License at
#
#        http://www.apache.org/licenses/LICENSE-2.0
#
#    Unless required by applicable law or agreed to in writing, software
#    distributed under the License is distributed on an "AS IS" BASIS,
#    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#    See the License for the specific language governing permissions and
#    limitations under the License.
#
# -*- coding: utf-8 -*-

# Author: whiteblue
# Created : 16/8/29

from __future__ import (absolute_import, division, print_function)

import time
import json

from ansible.utils.unicode import to_bytes
from ansible.plugins.callback import CallbackBase

__metaclass__ = type

FAILED = -1
UNREACHABLE = -2
ASYNC_FAILED = -3
OK = 0


class DemoCallbackModule(CallbackBase):
    """
    logs playbook results, per host
    """
    CALLBACK_VERSION = 2.0
    CALLBACK_TYPE = 'notification'
    CALLBACK_NAME = 'log_plays'
    CALLBACK_NEEDS_WHITELIST = True

    TIME_FORMAT = "%b %d %Y %H:%M:%S"
    MSG_FORMAT = "%(now)s - %(category)s - %(data)s\n\n"

    def __init__(self, step_callback, debug=False):
        super(DemoCallbackModule, self).__init__()
        self._debug = debug
        self._step_callback = step_callback

    def log(self, host, category, data):
        if type(data) == dict:
            if '_ansible_verbose_override' in data:
                data = 'omitted'
            else:
                data = data.copy()
                invocation = data.pop('invocation', None)
                data = json.dumps(data)
                if invocation is not None:
                    data = json.dumps(invocation) + " => %s " % data

        now = time.strftime(self.TIME_FORMAT, time.localtime())
        msg = to_bytes(self.MSG_FORMAT % dict(now=now, category=category, data=data))

        if self._debug:
            print(msg)

    def runner_on_failed(self, host, res, ignore_errors=False):
        self.log(host, FAILED, res)

    def runner_on_ok(self, host, res):
        self.log(host, OK, res)

    def runner_on_unreachable(self, host, res):
        self.log(host, UNREACHABLE, res)

    def runner_on_async_failed(self, host, res, jid):
        self.log(host, ASYNC_FAILED, res)
