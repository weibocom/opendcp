#!/bin/bash

# Author: whiteblue
# Time  : 2016/08/31


# check golint
command -v golint >/dev/null 2>&1 || { echo "not found lint, install.."; go get -u github.com/golang/lint/golint; }


# check code
ls |grep -v 'tests' |awk '{cmd="golint "$1;system(cmd)}'
