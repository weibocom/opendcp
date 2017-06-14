#!/usr/bin/env bash

echo "start building..."

cp hosts /etc/hosts

gopm build

./jupiter
