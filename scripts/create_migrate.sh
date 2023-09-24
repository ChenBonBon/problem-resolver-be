#!/bin/bash

if [ "$1" ]
then
   migrate create -ext sql -dir build/migrations -seq create_"$1"_table
else
   echo "migrate 名称不能为空"
   exit 0
fi
