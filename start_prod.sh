#!/usr/bin/env bash
[[ `systemctl status postgresql | awk '/Active/{print $2}'` == inactive ]] && sudo systemctl start postgresql
[[ `systemctl status redis | awk '/Active/{print $2}'` == inactive ]] && sudo systemctl start redis

go install
RUN_MODE=production webserver