#!/usr/bin/env bash
path=$(cd `dirname $0`;pwd)
nohup ${path}/birthdayreminder > ${path}/birthdayreminder.log 2>&1 &