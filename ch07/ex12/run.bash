#!bin/bash
go build main.go
go build ../fetch/fetch.go
./main & 
# access to http://localhost:8000/list

