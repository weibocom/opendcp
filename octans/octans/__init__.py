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
