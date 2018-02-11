#!/bin/bash
go build netcat/netcat3.go
go build reverb/reverb1.go
./reverb1 &
./netcat3
