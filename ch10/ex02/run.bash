#!/bin/bash
go build main.go
go build unarchive/unarchive.go
go build tar/tar.go
go build zip/zip.go
tar cvf sample.tar sample
zip sample.zip sample
./main
