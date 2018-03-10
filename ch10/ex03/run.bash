#!bin/bash
go build ./fetch/main.go
./main http://gopl.io/ch1/helloworld?go-get=1 | grep go-import
