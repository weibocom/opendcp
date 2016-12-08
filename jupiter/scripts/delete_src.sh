#!/bin/sh

FOLDERS="Dockerfile     build.sh       dao            interceptor    lastupdate.tmp provider   sql       vendor
Godeps         conf           docs           main.go        response       scripts        ssh
README.md      controllers    future         jupiter.iml    models         routers        service        tests logstore"
for FOLDER in $FOLDERS; do
    rm -rf $FOLDER
done
