#!/usr/bin/env bash 

[[ `systemctl status postgresql | awk '/Active/{print $2}'` == inactive ]] && sudo systemctl start postgresql

for val in \
    WS_DB_USER
do
    if [ -z ${!val} ]; then
        printf "error: $val not defined\n" >&2
        MISSING_VAR=true
    fi
done

if [ $MISSING_VAR ]; then
    exit 1
fi

echo Creating user ${WS_DB_USER}...
createuser $WS_DB_USER
# createuser --interactive --pwprompt
