#!/bin/bash
go build gopl.io/ch8/netcat1
go run main.go -port 8010 &
./netcat1
