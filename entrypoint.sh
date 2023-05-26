#!/bin/sh

sh -c '
    echo Envs are:
    if [ -z $INPUT_VERBOSE ] 
    then
        VERBOSE=""
    else
        VERBOSE="-verbose"
    fi
    if [ -z $INPUT_VERSION ] 
    then
        VERSION=""
    else
        VERSION="-version"
    fi
    echo command is: migrate -path $INPUT_PATH -database $INPUT_DATABASE -prefetch $INPUT_PREFETCH -lock-timeout $INPUT_LOCK_TIMEOUT $VERBOSE $VERSION $INPUT_COMMAND
    migrate -path $INPUT_PATH -database $INPUT_DATABASE -prefetch $INPUT_PREFETCH -lock-timeout $INPUT_LOCK_TIMEOUT $VERBOSE $VERSION $INPUT_COMMAND
'
