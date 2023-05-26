#!/bin/sh

sh -c '
    echo Envs are:
    if [ -z $INPUT_VERBOSE ] 
    then
        VERBOSE="fuck"
    else
        VERBOSE="-verbose"
    fi
    echo $INPUT_PATH
    echo $INPUT_DATABASE
    echo $INPUT_VERBOSE
    echo verbose flag will be $VERBOSE;
    echo $INPUT_COMMAND
    echo Envs end;
    echo command is: migrate -path $INPUT_PATH -database $INPUT_DATABASE $VERBOSE $INPUT_COMMAND
    migrate -path $INPUT_PATH -database $INPUT_DATABASE $VERBOSE $INPUT_COMMAND
'
