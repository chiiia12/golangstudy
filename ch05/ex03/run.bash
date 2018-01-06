#!/bin/bash
go build ../fetch/fetch.go
go build main.go
./fetch https://golang.org | ./main
