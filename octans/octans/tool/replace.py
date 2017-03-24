from sys import argv
with open('/data/octans/config.yml','r') as r:
    lines=r.readlines()
with open('/data/octans/config.yml','w') as w:
    for l in lines:
       if l.startswith('mysql'):
           w.write(l.replace('mysql://root:@127.0.0.1/octans?charset=utf8',argv[1]))
           continue
       if l.startswith('get_key_url'):
           w.write(l.replace('http://127.0.0.1:8888/v1/instance/sshkey/',argv[2]))
           continue
       w.write(l)

with open('/data/octans/octans/tool/autopushssh.py','r') as r:
    lines=r.readlines()
with open('/data/octans/octans/tool/autopushssh.py','w') as w:
    for l in lines:
       if l.startswith('    r = requests.put'):
           w.write(l.replace('127.0.0.1:7070', argv[3]))
           continue
       if l.startswith('    res = requests.put'):
           w.write(l.replace('127.0.0.1:7070', argv[3]))
           continue
       w.write(l)
