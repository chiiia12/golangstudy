#!/bin/bash
go build ../fetch/fetch.go
./fetch http://www.w3.org/TR/2006/REC-xml11-20060816 > w3.xml
go build main.go 
./main < w3.xml > w3.txt
