#!/bin/bash

service nginx start
service php-fpm start
service redis start

while true; do
        sleep 100
done
