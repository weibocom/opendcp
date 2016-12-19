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

import sys
import time

from collections import namedtuple

from ansible.parsing.dataloader import DataLoader
from ansible.vars import VariableManager
from ansible.inventory import Inventory
from ansible.playbook.play import Play
from ansible.inventory.host import Host
from ansible.inventory.group import Group
from ansible.executor.task_queue_manager import TaskQueueManager
from ansible.executor.task_executor import TaskExecutor
from ansible.executor.playbook_executor import PlaybookExecutor

from multiprocessing import Process, Pool

from threading import Thread

Options = namedtuple('Options',
                     ['listtags', 'listtasks', 'listhosts', 'syntax', 'connection', 'module_path', 'forks',
                      'remote_user', 'ansible_ssh_pass', 'ssh_common_args', 'ssh_extra_args', 'sftp_extra_args',
                      'scp_extra_args', 'become', 'become_method', 'become_user', 'verbosity', 'check'])

Loader = DataLoader()



def run():
    # bind sys-out
    origin = sys.stdout
    f = open('ansible.log', 'w')
    sys.stdout = f

    host = ["123.207.136.186"]
    variable_manager = VariableManager()
    loader = DataLoader()
    inventory = Inventory(loader=loader, variable_manager=variable_manager, host_list=host)

    # other options
    options = Options(listtags=False, listtasks=False, listhosts=False, syntax=False, connection='ssh',
                      module_path=".", forks=5,
                      remote_user='whiteblue', ansible_ssh_pass="", ssh_common_args=None,
                      ssh_extra_args=None,
                      sftp_extra_args=None, scp_extra_args=None, become=False, become_method=None,
                      become_user="whiteblue",
                      verbosity=None, check=False)

    executor = PlaybookExecutor(playbooks=["test.yml"], inventory=inventory,
                                variable_manager=variable_manager,
                                loader=Loader, options=options, passwords={})

    results = executor.run()

    print(type(results))
    print("result:" + str(results))


if __name__ == "__main__":
    run()
