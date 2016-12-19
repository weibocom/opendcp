#!/bin/bash

gopm build -o main && \
./main orm syncdb -v
#./main orm syncdb -force=0 -v

rm -f main