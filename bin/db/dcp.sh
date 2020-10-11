#!/usr/bin/env bash 

[[ `systemctl status postgresql | awk '/Active/{print $2}'` == inactive ]] && sudo systemctl start postgresql

./drop_db.sh
./create_db.sh
./populate_db.sh
