package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "./tar"
	_ "./zip"
	"./unarchive"
	"fmt"
)

func main() {
	unarchiver, err := unarchive.OpenUnArchiver("zip", "./sample.zip")
	if err != nil {
		fmt.Println(err)
	}
	unarchiver.UnArchive()

}
