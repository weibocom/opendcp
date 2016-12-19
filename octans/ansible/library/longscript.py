# !/usr/bin/env python
# -*- coding: utf-8 -*-

# Author: WhiteBlue
# Time  : 2016/07/28


import os
import stat
from uuid import uuid1
import datetime
import subprocess
from Queue import Queue

from ansible.module_utils.basic import AnsibleModule



DOCUMENTATION = '''
---
module: longscript
version_added: 0.1
short_description: run a long shell script in content
description:
     - run a shell script in content (by write local tmp file)
options:
  content:
    description:
      - the script content.
    required: true
requirements:
  - python >= 2.6

author: "WhiteBlue"

'''

EXAMPLES = '''
  - name: script...
    longscript: content="echo '233'"
'''

_TMP_DIR = "/tmp/"

_CMD_SEND = 0
_CMD_STOP = 1


def _write_file(content):
    tmp_file = _TMP_DIR + str(uuid1())

    with open(tmp_file, 'w') as f:
        f.write(content)

    os.chmod(tmp_file, stat.S_IRWXU)

    return tmp_file


def handle_log(tmp_arr, queue):
    while True:
        cmd, content = queue.get()
        if cmd == _CMD_SEND:
            pass
        elif cmd == _CMD_STOP:
            break
        else:
            raise Exception("unsupported command")


def main():
    module = AnsibleModule(
        argument_spec=dict(
            content=dict(required=True),
            remove=dict(),
            callback_url=dict(),
            sync_time=dict()
        )
    )

    content = module.params['content']

    tmp_file = None

    try:
        start_time = datetime.datetime.now()

        try:
            tmp_file = _write_file(content)
        except Exception as e:
            module.fail_json(msg='write script error, error: {}'.format(str(e)))
            return

        queue = Queue()

        ret_code = 0
        stdout = None
        try:
            p = subprocess.Popen(['/bin/bash', tmp_file], env=os.environ, shell=False, stdout=subprocess.PIPE,
                                 stderr=subprocess.PIPE, stdin=subprocess.PIPE)
            
            while True:
                buff = p.stdout.readline()
                
                if buff == '' and p.poll() != None:
                    queue.put((_CMD_STOP, None,))
                    break
                else:
                    queue.put((_CMD_SEND, buff))

            p.wait()
            (stdout, stderr) = p.communicate()
            if not isinstance(stderr, (bytes, unicode)):
                stderr = stderr.read()
            if not isinstance(stdout, (bytes, unicode)):
                stdout = stdout.read()

            ret_code = p.returncode

            if ret_code != 0:
                raise Exception("script return code {}, error: {}".format(ret_code, stderr))

        except Exception as e:
            module.fail_json(msg='run script error, error: {}'.format(str(e)))

        end_time = datetime.datetime.now()

        delta = end_time - start_time

        module.exit_json(
            stdout=stdout,
            rc=ret_code,
            start=str(start_time),
            end=str(end_time),
            delta=str(delta),
            changed=True,
        )

    except Exception as e:
        module.fail_json(msg='unexpected error, error: {}'.format(str(e)))
    finally:
        if tmp_file is not None:
            os.remove(tmp_file)


main()
