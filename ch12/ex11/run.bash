#!/bin/bash
go build ../../ch05/fetch/fetch.go
go run params.go &
./fetch 'http://localhost:12345/search'
