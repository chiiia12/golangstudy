#!/bin/bash
rm -rf ./out
mkdir out
go build main.go
go build unarchive/unarchive.go
go build tar/tar.go
go build zip/zip.go
tar cvf sample.tar sample
zip -r sample.zip sample
./main -type zip -filename ./sample.zip
#./main -type tar -filename ./sample.tar
