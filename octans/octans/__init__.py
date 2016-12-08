# !/usr/bin/env python
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
