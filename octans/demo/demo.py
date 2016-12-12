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
from collections import namedtuple

from ansible.parsing.dataloader import DataLoader
from ansible.vars import VariableManager
from ansible.inventory import Inventory
from ansible.playbook.play import Play
from ansible.inventory.host import Host
from ansible.inventory.group import Group
from ansible.executor.task_queue_manager import TaskQueueManager


Options = namedtuple('Options',
                     ['listtags', 'listtasks', 'listhosts', 'syntax', 'connection', 'module_path', 'forks',
                      'remote_user', 'ssh_common_args', 'ssh_extra_args', 'sftp_extra_args',
                      'scp_extra_args', 'become', 'become_method', 'become_user', 'verbosity', 'check'])

loader = DataLoader()


# demo by whiteblue
def handle():
    host1 = Host("127.0.0.1")

    host1.vars = dict(ansible_port=22,
                      ansible_user="root",
                      ansible_ssh_private_key_file="../tmp/127.0.0.1")

    # none variable
    variable_manager = VariableManager()

    g = Group("group_wb")
    g.add_host(host1)

    # target ip list
    inventory = Inventory(loader=loader, variable_manager=variable_manager)

    inventory.add_group(g)

    # other options

    options = Options(listtags=False, listtasks=False, listhosts=False, syntax=False, connection='ssh',
                      module_path=None, forks=1,
                      remote_user='root', ssh_common_args=None,
                      ssh_extra_args=None,
                      sftp_extra_args=None, scp_extra_args=None, become=False, become_method=None,
                      become_user="root",
                      verbosity=None, check=False)

    tqm = TaskQueueManager(
        inventory=inventory,
        variable_manager=variable_manager,
        loader=loader,
        options=options,
        passwords=None
    )

    # test tasks
    task0 = dict(action=dict(module='shell', args='sleep 6'))
    task1 = dict(action=dict(module='shell', args='ls'))
    task2 = dict(action=dict(module='shell', args='echo "2333"'))

    play_source = dict(
        name="Test Play",
        hosts="group_wb",
        gather_facts='no',
        tasks=[
            task0, task1, task2
        ]
    )

    play = Play().load(play_source, variable_manager=variable_manager, loader=loader)

    try:
        result = tqm.run(play)

        print("Result code: " + str(result))
        print("Type: " + str(type(result)))

    # close queue manager
    finally:
        if tqm is not None:
            tqm.cleanup()


if __name__ == "__main__":
    handle()
