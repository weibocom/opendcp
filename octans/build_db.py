#!/usr/bin/env python
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
