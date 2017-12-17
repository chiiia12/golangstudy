#!/bin/bash
#show
go run main.go -command show
#create
go run main.go -command create -token $GITHUB_TOKEN -title hoge -body piyo  
#update
go run main.go -command update -token $GITHUB_TOKEN -issue 1 -title change -body chang
#close
go run main.go -command close -token $GITHUB_TOKEN -issue 1

