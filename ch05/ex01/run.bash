#/bin/bash
go build fetch.go
go build main.go
./fetch https://golang.org | ./main
