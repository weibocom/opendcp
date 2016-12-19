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

import datetime

from sqlalchemy import and_
from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker

from octans.task.model import Task, Node, Log
from octans.logger import LogManager

Logger = LogManager.get_logger("SyncCallbackModule")

# Base class for Database operation
class BaseService:
    def __init__(self, config,pool_size,pool_recycle):
        self.engine = create_engine(config, pool_size=pool_size,pool_recycle=pool_recycle)
        self.session = sessionmaker(bind=self.engine)

    def _start_session(self):
        return self.session()

# Octans Database operation
class TaskService(BaseService):
    
    #status for task and node
    STATUS_INIT = 0
    STATUS_RUNNING = 1
    STATUS_SUCCESS = 2
    STATUS_FAILED = 3
    STATUS_STOPPED = 4
    STATUS_PartlySuccess = 5

    #record task to db (when starting a new task)
    def new_task(self, data):
        # type: (object) -> object
        session = self._start_session()
        try:
            now_time = datetime.datetime.now()

            obj = Task(**data)
            obj.status = TaskService.STATUS_INIT
            obj.create_time = now_time
            obj.update_time = now_time
            session.add(obj)
            session.commit()
            return obj.id
        except Exception as err:
            session.rollback()
            raise err
        finally:
            session.close()

    def new_node(self, task_id, ip):
        session = self._start_session()
        try:
            obj = Node()
            obj.task_id = task_id
            obj.ip = ip
            obj.status = TaskService.STATUS_INIT

            session.add(obj)
            session.commit()
            return obj.id
        except Exception as err:
            session.rollback()
            raise err
        finally:
            session.close()

    def add_log(self, global_id, source, task_uuid, task_status, create_time, end_time, data=None, host=""):
        session = self._start_session()
        try:
            obj = Log()
            obj.host = host
            obj.global_id = global_id
            obj.source = source
            obj.task_uuid = task_uuid
            obj.task_status = task_status
            obj.create_time = create_time
            obj.end_time = end_time
            if data is not None:
                obj.log = data
            session.add(obj)
            session.commit()
            return obj.id
        except Exception as err:
            session.rollback()
            raise err
        finally:
            session.close()

    def update_log(self,host, global_id, task_uuid, task_status, end_time, data=None):
        session = self._start_session()
        try:
            obj = session.query(Log).filter(Log.task_uuid==task_uuid,Log.global_id==global_id, Log.host==host).first() 
            obj.end_time = end_time
            obj.task_status = task_status
            obj.end_time = end_time
            obj.host = host
            if data is not None:
                obj.log = data
            session.commit()
            return obj.id
        except Exception as err:
            session.rollback()
            raise err
        finally:
            session.close()


    def check_task(self, task_id):
        session = self._start_session()
        try:
            query = session.query(Node).filter_by(task_id=task_id)
            ret = query.all()
            return ret
        finally:
            session.close()

    def get_task_by_id(self, task_id):
        session = self._start_session()
        try:
            obj = session.query(Task).filter_by(id=task_id).first()

            return obj
        finally:
            session.close()

    def get_log_by_globalid_source_host(self, global_id, source, host=None):
        session = self._start_session()  
        try:
            if host == None:
                obj = session.query(Log).filter_by(global_id=global_id, source=source).all()
                return obj
            else:
                obj = session.query(Log).filter_by(global_id=global_id, source=source, host=host).all()
                return obj
        finally:
            session.close()

    def get_task_by_name(self, task_name):
        session = self._start_session()
        try:
            obj = session.query(Task).filter_by(name=task_name).first()

            return obj
        finally:
            session.close()
            
    def get_node_by_id(self, node_id):
        session = self._start_session()
        try:
            obj = session.query(Node).filter_by(id=node_id).first()

            return obj
        finally:
            session.close()

    def update_node(self, node_id, status=None, log=None):
        session = self._start_session()
      
        try:
            obj = session.query(Node).filter_by(id=node_id).first()
            if obj is None:
                raise Exception("node not found...")

            obj.update_time = datetime.datetime.now()

            if status is not None:
                obj.status = status

            if log is not None:
                obj.log = log
           
            session.add(obj)
            session.commit()
        except Exception as err:
            session.rollback()
            raise err
        finally:
            session.close()

    def update_task(self, task_id, status=None, err=None):
        session = self._start_session()

        try:
            obj = session.query(Task).filter_by(id=task_id).first()
            if obj is None:
                raise Exception("node not found...")

            obj.update_time = datetime.datetime.now()

            if status is not None:
                obj.status = status

            if err is not None:
                obj.err = err

            session.add(obj)
            session.commit()
        except Exception as err:
            session.rollback()
            raise err
        finally:
            session.close()

