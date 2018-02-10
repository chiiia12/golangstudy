#!/bin/bash
kill `lsof -i:8000`
kill `lsof -i:8010`
kill `lsof -i:8020`
kill `lsof -i:8030`

go build ./main.go
go build ./clockwall.go

TZ=US/Eastern ./main -port 8010 &
TZ=Asiz/Tokyo ./main -port 8020 &
TZ=Europe/London ./main -port 8030&

./clockwall NewYork=localhost:8010 Tokyo=localhost:8020 London=localhost:8030
