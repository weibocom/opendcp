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

# Author: whiteblue
# Created : 16/8/16

from sqlalchemy.ext.declarative import declarative_base

from sqlalchemy import Table, Column, Integer, String, DateTime, Boolean, TEXT, Enum, ForeignKey, MetaData

#from octans.task import TABLE_PREFIX
from sqlalchemy import create_engine
Base = declarative_base()
TABLE_PREFIX = 'chan_'

class Task(Base):
    __tablename__ = TABLE_PREFIX + 'task'

    id = Column(Integer, primary_key=True)
    name = Column(String(255), unique=True)

    status = Column(Integer)

    err = Column(TEXT)

    create_time = Column(DateTime)
    update_time = Column(DateTime)

    def __repr__(self):
        return "<Task(id='{}', nodes='{}', status='{}'>".format(
            self.id, self.nodes, self.status)


class Node(Base):
    __tablename__ = TABLE_PREFIX + 'node'

    id = Column(Integer, primary_key=True)
    ip = Column(String(255), index=True)

    status = Column(Integer)

    task_id = Column(Integer, ForeignKey(Task.__tablename__ + ".id", onupdate="CASCADE", ondelete="CASCADE"))

    log = Column(TEXT)

    def __repr__(self):
        return "<Node(id='{}', ip='{}', status='{}'>".format(
            self.id, self.ip, self.status)

class Log(Base):
    __tablename__ = TABLE_PREFIX + 'log'

    id = Column(Integer, primary_key=True)
    host = Column(String(255))
    global_id = Column(String(255))
    source = Column(String(255))
    log = Column(TEXT)
    create_time = Column(String(255))
    end_time = Column(String(255))
    task_uuid = Column(String(255))
    task_status = Column(String(255))

    def __repr__(self):
        return "<Log(id='{}', host='{}', global_id='{}', source='{}',log='{}'>".format(
            self.id, self.host, self.global_id, self.source, self.log)

if __name__ == '__main__':
    engine = create_engine("mysql://root:roottoor@127.0.0.1/pytest", pool_size=50)
    Base.metadata.drop_all(engine)
    Base.metadata.create_all(engine)
