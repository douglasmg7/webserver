#!/usr/bin/env bash 

[[ `systemctl status postgresql | awk '/Active/{print $2}'` == inactive ]] && sudo systemctl start postgresql

for val in WS_DB_NAME
do
    if [ -z ${!val} ]; then
        printf "error: $val not defined\n" >&2
        MISSING_VAR=true
    fi
done

if [ $MISSING_VAR ]; then
    exit 1
fi

echo Droppping db $WS_DB_NAME
dropdb $WS_DB_NAME
