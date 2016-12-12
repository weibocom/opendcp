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

from flask import Flask

from octans.task.service import TaskService
from octans.worker.executor import Executor
from logger import LogManager
from config import Config

Conf = Config()
Conf.load("config.yml")
App = Flask('Octans')
Service = TaskService(Conf.get("mysql"),Conf.get("pool_size"),pool_recycle=7200)
Worker = Executor(Service)


def start_app():
    Worker.start()


import octans.api
