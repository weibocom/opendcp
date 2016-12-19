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

# Author: WhiteBlue, Fuyuhui
# Time  : 2016/07/20

import json
import os
import stat
from collections import namedtuple

import requests
from ansible.executor.task_queue_manager import TaskQueueManager
from ansible.inventory import Inventory
from ansible.inventory.group import Group
from ansible.inventory.host import Host
from ansible.parsing.dataloader import DataLoader
from ansible.playbook.play import Play
from ansible.vars import VariableManager

from octans import Service, Conf
from octans.ansible.ansible_callback import SyncCallbackModule
from octans.logger import LogManager
from octans.worker.task import Task
from _ast import IsNot

_Options = namedtuple('Options',
                      ['connection', 'module_path', 'forks', 'timeout', 'remote_user', 'private_key_file',
                       'ssh_common_args', 'ssh_extra_args', 'sftp_extra_args', 'scp_extra_args',
                       'become', 'become_method', 'become_user', 'verbosity', 'check'])

_Loader = DataLoader()

_Loader.set_basedir("./ansible")

Logger = LogManager.get_logger("AnsibleTask")


def _get_ssh_key(ip):
    """
    get ssh_private_key content from api

    :param ip: target ip
    :return: content in str
    """
    ret = requests.get(Conf.get("get_key_url") + str(ip), timeout=5)
   
    if ret.ok:
        if ret.content is None:
            raise Exception("Blank SSH key returened")
        # back_json = json.loads(ret.content)
        # key = back_json["data"]["key"]
        return ret.content
    else:
        raise Exception("connection failed...")
        # with open("/Users/whiteblue/.ssh/id_rsa") as f:
        #     return f.read()


def _write_ssh_key(filename, key_content):
    """
    write ssh_private_key file

    :param filename: filename
    :param key_content: content
    :return:
    """
    key_file = "tmp/" + filename

    with open(key_file, 'w') as f:
        f.write(key_content)
    
    os.chmod(key_file, stat.S_IRUSR | stat.S_IWUSR)
    return key_file


def _rm_tmp_key(key_files):
    try:
        for path in key_files:
            os.remove(path)
    except Exception as e:
        Logger.error("rm tmp ssh_key file failed, error:{}".format(str(e)))


class AnsibleTask(Task):
    """
    run 'ansible' playbook as a Task
    """


    def __init__(self, task_id, name, hosts, tasks,tasktype, user, global_id, source, result=None, forks=5,params=None):

        """
        init task and make playbook instance

        :param task_id: name used in playbook and host group
        :param hosts: target hosts(list)
        :param tasks: 'ansible' tasks(list)
        :param params: params used in each ost
        """
        Task.__init__(self, id=task_id)

        if user is None:
            user = "root"

        self.hosts = hosts
        self.params = params
        self.task_id = task_id
        self.tasks = tasks
        self.tasktype = tasktype
        self.log = []
        self.log_iter = 0
        self.result = result
        self.forks = forks
        self.name = name
        self.user = user.strip()
        self._node_map = dict()
        self.global_id = global_id
        self.source = source

    def _step_callback(self, ip, code, data=None):
        try:
            
            logdata=[]
            node=Service.get_node_by_id(node_id=self._node_map[ip]);


            if (node is not None) and (node.log is not None):
               
                temp=eval(node.log)
                Logger.debug("node.log:"+node.log);
                if isinstance(temp,  list):
                  
                    logdata=temp;
                elif isinstance(temp,  dict):
                   
                    logdata=[temp]
                
           
            logdata=json.dumps(logdata)
            if code < 0:
              
                Service.update_node(node_id=self._node_map[ip], status=Service.STATUS_FAILED, log=logdata)
            elif code==0:
                Service.update_node(node_id=self._node_map[ip], status=Service.STATUS_SUCCESS, log=logdata)
            else:
                Service.update_node(node_id=self._node_map[ip], status=Service.STATUS_RUNNING, log=logdata)
        except Exception as e:
            Logger.info("step callback falied, ip: {}, error: {}, global id: {}".format(ip, e.message, self.global_id))

    def run(self):
        # insert node
        for ip in self.hosts:
            self._node_map[ip] = Service.new_node(self.task_id, ip)

        variable_manager = VariableManager()

        Logger.debug("start write ssh_key for task: {} global_id : {}".format(self.task_id, self.global_id))

        key_files = []

        group = Group(self.task_id)

        for h in self.hosts:
            
            # get ssh_key content
            key_content = _get_ssh_key(h)

            Logger.debug("read ssh_key for host: {} global_id: {}".format(h, self.global_id))

            # write ssh private key
            key_path = _write_ssh_key(h, key_content)

            #key_path="./tmp/97"
            Logger.debug("write ssh_key for host: {} global_id: {}".format(h, self.global_id))

            host_vars = dict(ansible_port=22,
                             ansible_user=self.user,
                             ansible_ssh_private_key_file="./" + key_path)

            Logger.debug("key_path: {} global_id: {}".format(key_path, self.global_id))

            key_files.append(key_path)

            host = Host(h)

            host.vars = host_vars

            group.add_host(host)

        # add params to each host
        if self.params is not None and isinstance(self.params, dict):
            for h in group.hosts:
                for key in self.params.keys():
                    variable_manager.set_host_variable(h, key, self.params[key])

        Logger.debug("success write ssh_key for task: {} global_id: {}".format(self.task_id, self.global_id))

        # other options
        ssh_args = '-oControlMaster=auto -oControlPersist=60s -oStrictHostKeyChecking=no'
        options = _Options(connection='ssh', module_path='./ansible/library',
                           forks=self.forks,
                           timeout=10,
                           remote_user=None,
                           private_key_file=None,
                           ssh_common_args=ssh_args,
                           ssh_extra_args=None,
                           sftp_extra_args=None, scp_extra_args=None, become=None,
                           become_method=None,
                           become_user=None, verbosity=None, check=False)

        if self.tasktype=="ansible_task":
            Logger.debug("ansible tasks set*******************  global_id: {}".format(self.global_id))
            play_source = dict(
                name=self.task_id,
                hosts=self.task_id,
                gather_facts='yes',
                tasks=self.tasks
            )
        else:
            
            Logger.debug("ansible role set******************* global_id: {}".format(self.global_id))
            play_source = dict(
                name=self.task_id,
                hosts=self.task_id,
                gather_facts='yes',
                roles=self.tasks
            )
                

        Logger.debug("start load play for task: {} global_id: {}".format(self.task_id, self.global_id))

        # make playbook
        playbook = Play().load(play_source, variable_manager=variable_manager, loader=_Loader)

        inventory = Inventory(loader=_Loader, variable_manager=variable_manager)

        inventory.add_group(group)

        call_back = SyncCallbackModule(debug=True, step_callback=self._step_callback, global_id=self.global_id, source=self.source, tag_hosts=self.hosts)

        Logger.debug("success load play for task: {} global_id: {}".format(self.task_id, self.global_id))

        # task queue
        tqm = TaskQueueManager(
            inventory=inventory,
            variable_manager=variable_manager,
            loader=_Loader,
            options=options,
            passwords=None,
            stdout_callback=call_back
        )

        try:
            back = tqm.run(playbook)

            Logger.info("back: {} global_id : {}".format(str(back), self.global_id))

            if back != 0:
                raise Exception("playbook run failed")

            return back
        finally:
            if tqm is not None:
                tqm.cleanup()
                _rm_tmp_key(key_files)

    def failed(self, error):
        err_json = dict(msg=str(error))
        
        #check whether success node exist
        node_list = Service.check_task(task_id=str(self.task_id))
        successflag =False
        for node in node_list:
                if node.status==2:
                    successflag=True
                    break
        # update task
        if successflag:
            Service.update_task(task_id=self.task_id, status=Service.STATUS_PartlySuccess, err=json.dumps(err_json))
        else:
           
            Service.update_task(task_id=self.task_id, status=Service.STATUS_FAILED, err=json.dumps(err_json))

        # update nodes in task
#        for ip in self.hosts:
#            Service.update_node(node_id=self._node_map[ip], status=Service.STATUS_FAILED)

    def success(self, result):
        Service.update_task(task_id=self.task_id, status=Service.STATUS_SUCCESS)

        for ip in self.hosts:
            Service.update_node(node_id=self._node_map[ip], status=Service.STATUS_SUCCESS)
