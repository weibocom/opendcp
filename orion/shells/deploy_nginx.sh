#!/bin/sh

# get ip address
IP=`ifconfig eth0 | grep inet | awk '{print $2}'`
echo "IP is $IP"

# run role
echo "Deploy nginx on $IP ..."
NOW=`date +"%Y%m%d-%H%M%S"`
curl -l -H "Content-type: application/json" -X POST \
    -d  "{\"tasks\": [\"hubble-nginx\"], \"name\": \"$IP_$NOW\", \"fork_num\":5, \"tasktype\": \"ansible_role\", \"nodes\": [\"$IP\"], \"user\": \"root\"}" \
    http://{{octans_host}}:8082/api/parallel_run
