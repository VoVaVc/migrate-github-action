#!/bin/sh

set -eu

sh -c "
    echo $*
    migrate -path $1 -database $2 -verbose $3 $4
"