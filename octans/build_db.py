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
# Created : 16/8/16



from sqlalchemy import create_engine

from channel import Conf
from channel.task.model import Task, Node, Log


'''
database init
'''

engine = create_engine(Conf.get("mysql"), pool_size=50)

bases = [Task, Node, Log]
for base in bases:
    base.metadata.create_all(engine)
