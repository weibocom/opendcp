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
# Time  : 2016/07/26
import os
from multiprocessing import Process
from threading import Thread

import signal

import sys

from octans.logger import LogManager
from octans.worker.task import Task
from multiprocessing import Queue



'''
daemon thread to run task
'''
_CMD_NEWTASK = 0
_CMD_STOPTASK = 1
_CMD_CLEARTASK = 2

Logger = LogManager.get_logger("Executor")

_output_null = open("/dev/null", "w")


def _find_child_process(pid):
    pid_str = str(pid)
    sh = "ps -e -o pid,ppid | awk '{ if($1!=" + pid_str + "&&$2==" + pid_str + "){printf $1\",\"}}'"
    t = os.popen(sh)
    child_arr = t.read().split(",")
    ret = []
    for child in child_arr:
        if child.strip() is not "" and len(child) != 0:
            ret.append(child)
    return ret


def _clear_remain_process(pid):
    """
    kill the child process that Anisble fork
    :param pid: process pid
    :return: none
    """
    try:
        # make main process stop
        os.kill(pid, signal.SIGSTOP)
        kill_list = []
        find_list = [pid]
        count = 0
        while len(find_list) is not 0 and count < 20:
            t_pid = find_list.pop()
            ret_list = _find_child_process(t_pid)
            kill_list.extend(ret_list)
            find_list.extend(ret_list)
            count += 1

        Logger.debug("kill Ansible fork process: {}".format(str(kill_list)))
        for p in kill_list:
            try:
                os.kill(int(p), signal.SIGKILL)
            except Exception as e:
                Logger.error("kill process error, error:{}".format(str(e)))
    except Exception as e:
        Logger.error("clear process error, error:{}".format(str(e)))


class Executor:
    """
    a process executor with queue
    """

    def __init__(self, service):
        self._queue = Queue(-1)
        self._thread = None
        self._task_list = []
        self.service = service

    def start(self):
        """
        start the loop thread
        :return: none
        """
        t = Thread(target=self._loop_for_queue)
        t.setDaemon(True)
        t.start()

    def submit(self, task):
        """
        submit a new task to queue
        :param task: Task instance
        :return: none
        """
        if not isinstance(task, Task):
            raise AttributeError("'task' must instance of Type Task")
        self._queue.put((_CMD_NEWTASK, task,))

    def stop(self, task_id):
        """
        send SIGTERM to a process by task_id
        :param task_id: task id
        :return: none
        """
        self._queue.put((_CMD_STOPTASK, task_id,))

    def list(self):
        """
        list task in running
        :return:
        """
        ret = []
        for p in self._task_list:
            ret.append(p.name)
        return ret

    def _handle(self, task):
        """
           run in a new process
           :param task: task obj
           :return: none
           """
        try:
            sys.stdout = _output_null
            sys.stderr = _output_null

            try:
                Logger.info("task run start, task_id: {}".format(task.get_id()))
                ret = task.run()
                Logger.info("task run start-----------, task_ret: {}".format(ret))
                task.success(ret)
                Logger.info("task run success, task_id: {}".format(task.get_id()))
            except AttributeError as ae:
                Logger.info("ansible error see https://github.com/ansible/ansible/issues/14408 we pass --pengtao, ae: {}".format(str(ae)))
                pass
            except Exception as e:
                Logger.error("task run failed, task_id: {} ,error: {}".format(task.get_id(), str(e)))
                task.failed(e)
            finally:
                task.final()
        except Exception as out:
            Logger.error("task error, error: {}".format(out.message))
        finally:
            self._queue.put((_CMD_CLEARTASK, task.get_id(),))

    def _append_to_list(self, process):
        self._task_list.append(process)

    def _remove_from_list(self, process_name):
        for p in self._task_list:
            if p.name == process_name:
                self._task_list.remove(p)
                return p
        return None

    def _loop_for_queue(self):
        """
            run in a thread receive cmd & arg
        """
        while True:
            try:
                command, obj = self._queue.get()
                flag = True
                if command == _CMD_NEWTASK:
                    for p in self._task_list:
                        if obj.get_id() == p.name:
                            Logger.info("task already in list, task_id: {}".format(obj.get_id()))
                            flag = False
                    if flag:
                        Logger.debug("submit new task, task_id: {}".format(obj.get_id()))
                        p = Process(name=obj.get_id(), target=self._handle, args=(obj,))
                        p.daemon = False
                        p.start()
                        self._append_to_list(p)
                elif command == _CMD_STOPTASK:
                    # remove process from list
                    p = self._remove_from_list(obj)
                    if p is not None and isinstance(p, Process):
                        Logger.info("stop task, task_id: {}, pid: {}".format(p.name, p.pid))
                        # clear Ansible's fork process
                        os.kill(int(p.pid), signal.SIGINT)
                    else:
                        Logger.info("task not found, task_id: {}".format(obj))
                elif command == _CMD_CLEARTASK:
                    p = self._remove_from_list(obj)
                    if p is not None:
                        Logger.debug("task removed ,task_id: {}".format(obj))
                    else:
                        Logger.debug("task not removed ,task_id: {}".format(obj))
            except Exception as e:
                Logger.error("loop thread error, error: {}".format(str(e)))
