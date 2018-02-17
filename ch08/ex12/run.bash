#!/bin/bash
killall chat
go build chat.go
go build netcat3.go
./chat &
./netcat3
