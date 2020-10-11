#!/usr/bin/env bash 

[[ `systemctl status postgresql | awk '/Active/{print $2}'` == inactive ]] && sudo systemctl start postgresql

for val in \
    WS_DB_NAME \
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

echo Creating $WS_DB_NAME db...
createdb $WS_DB_NAME
psql -U $WS_DB_USER -d $WS_DB_NAME -f ./schema.sql
