#!/bin/sh

sh -c "echo $@"
sh -c "migrate -path $1 -database $2 -verbose $3 $4"