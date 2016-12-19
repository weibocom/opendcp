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

# Author: WhiteBlue,Fuyuhui
# Time  : 2016/07/26


from flask import request, jsonify, app

from octans import Worker, App, Service
from octans.ansible.ansible_task import AnsibleTask
import json
from octans.logger import LogManager

Logger = LogManager.get_logger("API")

#
#System http interface: run task, stop task and check task status
#
#Lack of param for invalid param type exception
class ParamErrorException(Exception):
    pass

#Error json format exception
class JsonEncodeException(Exception):
    pass

#Return api invoking with fail result
def return_failed(code=-1, message=""):
    return jsonify(dict(code=code, message=message))

#Return api invoking with success result
def return_success(code=0, content=None):
   
    if content is None:
        content = {}
    return jsonify(dict(code=code, content=content))

'''
check http params
    param    req_json: json object contains all http params
    param    key:    param name
    param    param_type:    param type in python
    param    default_value:    specify a default value if not fould in req_json
    param    allowNone:    raise exception if False and not fould in req_json
'''
def conform_param(req_json, key, param_type=None, default_value=None, allowNone=False):
    if req_json.has_key(key):
        value = req_json[key]
        if param_type is not None:
            if not isinstance(value, param_type):
                raise ParamErrorException("key " + key + " not instance of " + str(type))

        if param_type is int:
            try:
                value = int(value)
            except Exception:
                raise ParamErrorException("key " + key + " not instance of " + str(type))

        return value
    else:
        if default_value is None and not allowNone:
            raise ParamErrorException("key " + key + " not found")
        else:
            return default_value

#show introduce info to root request
@App.route('/')
def defaultpage():
    return jsonify(dict(name="Octans", version="0.1", info="an http ansible channel from Weibo"))



'''
Run system configuration task on specified ip
    param name: provide a new name for task, should not repeat with previous
    param nodes: target ip list
    param tasks: task list , element type is according to task type.For ansible_role, the element should be role name in string, 
                for ansible task, the element should be json format,for example:[{"action":{"args":"echo {{name}}","module":"shell"}}]
    param tasktype: currently supports two values: ansible_task , ansible_role
    param params: params used in ansible python api variable_manager
    param user: target system user to be used 
    param fork_num: thread count for parallel execution
'''
@App.route('/api/run', methods=['POST'])
def run_task():
    try:
        
        #read http params
        req_json = request.get_json(force=True, silent=True)
        if req_json is None:
            raise JsonEncodeException
        global_id = request.headers.get("X-CORRELATION-ID")
        if global_id is None:
            Logger.error("Missing X-CORRELATION-ID")
            return return_failed(-1, "X-CORRELATION-ID is empyt"), 400
        source = request.headers.get("X-SOURCE")
        if source is None:
            return return_failed(-1, "X-SOURCE is empyt"), 400
        Logger.debug("Run request json:"+json.dumps(req_json)+str(global_id))
        task_name = conform_param(req_json, "name", basestring)
        nodes = conform_param(req_json, "nodes", list)
        tasks = conform_param(req_json, "tasks", list)
        tasktype= conform_param(req_json, "tasktype", basestring,default_value="ansible_task")
        params = conform_param(req_json, "params", dict,{},True)
        user_name = conform_param(req_json, "user", basestring, allowNone=True)
        fork_num = conform_param(req_json, "fork_num", int, allowNone=True)
        # check task name duplicate
        task = Service.get_task_by_name(task_name)
        if task is not None:
            Logger.error("task name is duplicate:" + task_name)
            return return_failed(-1, "task name is duplicate"), 400
        task_id = Service.new_task({"name": task_name})

        #submit task
        Worker.submit(
            AnsibleTask(task_id=str(task_id), name=task_name, hosts=nodes, tasks=tasks,tasktype=tasktype, params=params, user=user_name,
                        forks=fork_num, global_id=global_id, source=source, result=""))

        return return_success(content={"id": task_id}), 200
    except JsonEncodeException as e:
        Logger.error("try run_task exception ---------@ ") 
        return return_failed(-1, "json encode error"), 400
    except ParamErrorException as e:
        Logger.error("try run_task exception --------------> %s" %(str(e))) 
        return return_failed(-1, "param error, error: " + e.message), 400
    except Exception as e:
        Logger.error("try run_task exception --------------> %s" %(str(e))) 
        return return_failed(-1, e.message), 500

'''
Stop task by id and name
    param    id: task id return after run request
    param    name: user specified task name in run request
'''
@App.route('/api/stop', methods=['POST'])
def stop_task():
    try:
        #read http params
        req_json = request.get_json(force=True, silent=True)
        if req_json is None:
            raise JsonEncodeException
        global_id = request.headers["X-CORRELATION-ID"]
        if global_id is None:
            return_failed(-1, "X-CORRELATION-ID is Empty"), 400
        source = request.headers["X-SOURCE"]
        if source is None:
            return_failed(-1, "X-SOURCE is Empty"), 400

        Logger.debug("Stop request json:"+json.dumps(req_json)+str(global_id))
        task_id = conform_param(req_json, "id", param_type=int, allowNone=True)
        task_name = conform_param(req_json, "name", param_type=basestring, allowNone=True)

        #stop task by id
        if task_id is None:
            if task_name is None:
                raise ParamErrorException("key task_id/task_name not found")
            else:
                task = Service.get_task_by_name(task_name)
                if task is None:
                    return return_failed(-1, "task not found"), 404
                task_id = task.id

        Worker.stop(str(task_id))

        # update task status
        Service.update_task(task_id, status=Service.STATUS_STOPPED)

        return return_success(), 200
    except JsonEncodeException:
        Logger.error("try stop_task exception --------------@") 
        return return_failed(-1, "json encode error"), 400
    except Exception as e:
        Logger.error("try stop_task exception --------------> %s" %(str(e))) 
        return return_failed(-1, e.message), 400

'''
Check status for specified request by id and name
    param    id: task id return after run request
    param    name: user specified task name in run request
'''
@App.route('/api/check', methods=['POST'])
def check_task():
    try:
        #read http params
        req_json = request.get_json(force=True, silent=True)
        if req_json is None:
            raise JsonEncodeException
        global_id = request.headers["X-CORRELATION-ID"]
        if global_id is None:
            return_failed(-1, "X-CORRELATION-ID is Empty"), 400
        source = request.headers["X-SOURCE"]
        if source is None:
            return_failed(-1, "X-SOURCE is Empty"), 400

        Logger.debug("Check request json:"+json.dumps(req_json)+str(global_id))
        task_id = conform_param(req_json, "id", param_type=int, allowNone=True)
        task_name = conform_param(req_json, "name", param_type=basestring, allowNone=True)

        #load task by id or name
        if task_id is None:
            if task_name is None:
                raise ParamErrorException("key task_id/task_name not found")
            else:
                task = Service.get_task_by_name(task_name)
                if task is None:
                    return return_failed(-1, "no task found for specified name:"+task_name), 404
        else:
            task = Service.get_task_by_id(task_id)
            if task is None:
                return return_failed(-1, "no task found for specified id:"+str(task_id)), 404
        
        node_list = Service.check_task(task_id=str(task.id))
        
        #return status data
        ret_node = []
        for node in node_list:
            ret_node.append(dict(
                ip=node.ip,
                status=node.status,
                #log=json.loads(node.log),
                log=node.log,
            ))

        ret_task = dict(
            id=task.id,
            status=task.status,
            err=task.err,
        )

        return return_success(content={
            "task": ret_task,
            "nodes": ret_node
        }), 200
    except JsonEncodeException:
        Logger.error("try check_task exception --------------@")
        return return_failed(-1, "json encode error"), 400
    except Exception as e:
        Logger.error("try check_task exception --------------> %s" %(str(e))) 
        return return_failed(-1, e.message), 500
'''
Check status for specified request by id and name
    param    id: task id return after run request
    param    name: user specified task name in run request
'''
@App.route('/api/getlog', methods=['POST'])
def getlog():
    try:
        #read http params
        req_json = request.get_json(force=True, silent=True)
        if req_json is None:
            raise JsonEncodeException
        global_id = request.headers["X-CORRELATION-ID"]
        if global_id is None:
            Logger.error("X-CORRELATION-ID is Empty")
            return_failed(-1, "X-CORRELATION-ID is Empty"), 400
        source = request.headers["X-SOURCE"]
        if source is None:
            Logger.error("X-CORRELATION-ID is Empty")
            return_failed(-1, "X-SOURCE is Empty"), 400
        Logger.debug("Check request json:"+json.dumps(req_json)+str(global_id))
        host = conform_param(req_json, "host", param_type=basestring, allowNone=True)
        source = conform_param(req_json, "source", param_type=basestring, allowNone=True)
        logs = None
        #load task by id or name
        if host is None:
            if source is None:
                raise ParamErrorException("source  not found")
            else:
                logs = Service.get_log_by_globalid_source_host(global_id, source)
                if logs is None:
                    return return_failed(-1, "no log found for specified name:"+source), 404
        else:
            logs = Service.get_log_by_globalid_source_host(global_id, source, host)
            if logs is None:
                return return_failed(-1, "no task found for specified global_id source host:"+str(global_id)), 404
        
        #return status data
        ret_log = []

            # ret_log.append(dict(
            #     global_id=log.global_id,
            #     source=log.source,
            #     log=json.loads(log.log),
            # ))
        for log in logs:
            if log.task_status == "failed":
                tmps = "global_id = %s, source = %s, create_time = %s, end_time = %s, host = %s, reuslt = %s " %(log.global_id, log.source, log.create_time, log.end_time, log.host,
 log.task_status)+"\n\t"
                redict = json.loads(log.log)
                if "results" not in redict.keys():
                    tmps += "message: "
                    if "msg" in redict.keys():
                        tmps += redict["msg"]
                    elif "stderr" in redict.keys():
                        tmps += redict["stderr"]
                    else:
                        tmps += "no error msg out"
                    tmps += "\n\t"
                    ret_log.append(tmps)
                    continue
                for i in redict["results"]:
                    if "msg" in redict.keys():
                    #if i["msg"] != "":
                        tmps += "message: "
                        tmps += i["msg"]
                        tmps += "\n\t"
                ret_log.append(tmps)
                continue
            if log.task_status == "unreacheable":
                tmps = "global_id = %s, source = %s, create_time = %s, end_time = %s, host = %s, reuslt = %s " %(log.global_id, log.source, log.create_time, log.end_time, log.host,
 log.task_status)+"\n\t"
                if "results" not in redict.keys():
                    tmps += "message: "
                    if "msg" in redict.keys():
                        tmps += redict["msg"]
                    elif "stderr" in redict.keys():
                        tmps += redict["stderr"]
                    else:
                        tmps += "no error msg out"
                    tmps += "\n\t"
                ret_log.append(tmps)
                continue
            ret_log.append("global_id = %s, source = %s, create_time = %s, end_time = %s, host = %s, reuslt = %s " %(log.global_id, log.source, log.create_time, log.end_time, log.host, log.task_status))
        return return_success(content={
            "log": ret_log
        }), 200
    except JsonEncodeException:
        Logger.error("try getlog exception --------------@")
        return return_failed(-1, "json encode error"), 400
    except Exception as e:
        Logger.error("try getlog exception --------------> %s" %(str(e))) 
        return return_failed(-1, e.message), 500
