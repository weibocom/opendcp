import requests
import json
import socket
import os
import hashlib
from Crypto.PublicKey import RSA

myname = socket.getfqdn(socket.gethostname())
myaddr = socket.gethostbyname(myname)

def get_resource():
    src = myname+myaddr
    m1 = hashlib.md5()
    m1.update(src)
    return m1.hexdigest()

def phydev_con():
    payload = {
                "InstanceId": get_resource(),
                "PrivateIpAddress":myaddr,
                "PublicIpAddress":myaddr,
                }

    headers ={'content-type': 'application/json',
            'Cache-Control':'no-cache',
            'App-Id':'tPa6NY5FDI9Fl3ie',
            'App-Key':'49sSs39RA614TLeVzLT2Z68Y4uY3BwZ7'}
    res = requests.put("http://127.0.0.1:7070/v1/instance/phydev",
                        data=json.dumps(payload),
                        headers=headers)
    if res.json().has_key("code"):
        if res.json()["code"] == 0:
            print(res.json())
            return
    raise Exception("phydev out control!")

def upload_sshkey():
    key = RSA.generate(2048)
    # print key.exportKey('PEM')

    payload = {
                "PublicKey": key.exportKey('OpenSSH'),
                "PrivateKey": key.exportKey('PEM')}
    headers ={'content-type': 'application/json',
            'Cache-Control':'no-cache',
            'App-Id':'tPa6NY5FDI9Fl3ie',
            'App-Key':'49sSs39RA614TLeVzLT2Z68Y4uY3BwZ7'}

    pubkey = key.publickey()
    if not os.path.exists("/root/.ssh"):
            os.makedirs("/root/.ssh")

    with open("/root/.ssh/authorized_keys", 'w') as content_file:
            # print pubkey.exportKey('OpenSSH')
            content_file.write(pubkey.exportKey('OpenSSH'))
    r = requests.put("http://127.0.0.1:7070/v1/instance/sshkey/%s" %(get_resource()),
                        data=json.dumps(payload),
                        headers=headers)
    if r.json().has_key("code"):
        if r.json()["code"] == 0:
            print(r.json())
            return
    raise Exception("Upload ssh key's Exception raised!")

upload_sshkey()
phydev_con()
