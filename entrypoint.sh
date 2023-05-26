#!/bin/sh

sh -c "
    echo Envs are:
    VERBOSE=$([ -z $INPUT_VERBOSE ] && echo "" || echo "-verbose");
    echo $INPUT_PATH
    echo $INPUT_DATABASE
    echo $INPUT_VERBOSE
    echo $INPUT_COMMAND
    echo Envs end;
    echo command is: migrate -path $INPUT_PATH -database $INPUT_DATABASE $VERBOSE $INPUT_COMMAND
    migrate -path $INPUT_PATH -database $INPUT_DATABASE $VERBOSE $INPUT_COMMAND
"
