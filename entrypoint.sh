#!/bin/sh

set -eu

export GITHUB="true"

sh -c "
    echo $*
    migrate -path $1 -database $2 -verbose $3 $4
"