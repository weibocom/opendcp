#!/bin/sh

FOLDERS="Dockerfile api build.sh    executor    main.go     routers     service     utils
Godeps      api.md      handler     models      style_check.sh  vendor
README.md   base_check.sh   controllers helper  scripts     tests       views"
for FOLDER in $FOLDERS; do
    rm -rf $FOLDER
done
