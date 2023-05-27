#!/bin/sh

sh -c '
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
    migrate -path $INPUT_PATH -database $INPUT_DATABASE -prefetch $INPUT_PREFETCH -lock-timeout $INPUT_LOCKTIMEOUT $VERBOSE $VERSION $INPUT_COMMAND
'
