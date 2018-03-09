package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "./tar"
	_ "./zip"
	"./unarchive"
)

func main() {
	unarchive.Read("zip")

}
