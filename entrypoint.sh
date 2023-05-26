#!/bin/sh

set -eu

sh -c "
    echo Envs are:
    verbose=''
    if [ ! -z '$INPUT_VERBOSE' -a '$INPUT_VERBOSE' != '' ]; then
        verbose='verbose'
    fi
    echo $INPUT_PATH
    echo $INPUT_DATABASE
    echo $INPUT_VERBOSE
    echo $INPUT_COMMAND
    echo Envs end;
    echo command is: migrate -path $INPUT_PATH -database $INPUT_DATABASE $verbose $INPUT_COMMAND
    migrate -path $INPUT_PATH -database $INPUT_DATABASE $verbose $INPUT_COMMAND
"
