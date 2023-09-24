#!/bin/bash

db_name=""
username=""
password=""
db_port=""
tables=""

if [ "$1" ]
then
    db_name="$1"
else
    echo "db_name 不能为空"
    exit 0
fi

if [ "$2" ]
then
    username="$2"
else
    echo "username 不能为空"
    exit 0
fi
  
if [ "$3" ]
then
    password="$3"
else
    echo "password 不能为空"
    exit 0
fi

if [ "$4" ]
then
    db_port="$4"
else
    echo "db_port 不能为空"
    exit 0
fi

if [ "$5" ]
then
    tables="$5"
fi

gentool -dsn "host=localhost user=$username password=$password dbname=$db_name port=$db_port sslmode=disable TimeZone=Asia/Shanghai" -db "postgres" -tables="$tables"