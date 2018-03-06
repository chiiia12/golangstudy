#!/bin/bash
go build main.go
tar cvf sample.tar sample
zip sample.zip sample
./main
