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

# Author: WhiteBlue
# Time  : 2016/07/20


class Task:
    """
    Base "task" Class
    """

    def __init__(self, id):
        self.id = id

    def get_id(self):
        return self.id

    def before(self):
        pass

    def run(self):
        pass

    def success(self, result):
        pass

    def failed(self, error):
        pass

    def final(self):
        pass
