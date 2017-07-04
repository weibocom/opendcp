#!/usr/bin/env bash

echo "start building..."

#cp hosts /etc/hosts

cp /etc/hosts /etc/hosts1
echo '127.0.0.1 controller' >> /etc/hosts1
cp /etc/hosts1 /etc/hosts
rm /etc/hosts1

gopm build


./jupiter
