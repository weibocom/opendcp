#!/bin/bash

# Author: whiteblue
# Time  : 2016/08/31


# check goimports
command -v goimports >/dev/null 2>&1 || { echo "Not found goimports, install.."; go get -u golang.org/x/tools/cmd/goimports; }


gopm build || { echo "Build failed ..."; exit 1; }

# base check
ERR_FILES=$(go tool vet .)
if [ -n "$ERR_FILES" ]; then
    echo "The following files has wrong code:"
    echo ${ERR_FILES}
    exit 1
fi


# check formatted
IMP_FILES=$(goimports -e -l .)

if [ -n "$IMP_FILES" ]; then
    echo "The following files are not properly formatted:"
    echo ${IMP_FILES}
    echo "Run auto format...."
    goimports -e -w ${IMP_FILES}
fi


echo "check passed"