package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/chiiia12/golangstudy/ch10/ex02/tar"
	_ "github.com/chiiia12/golangstudy/ch10/ex02/zip"
	"github.com/chiiia12/golangstudy/ch10/ex02/unarchive"
)

func main() {
	//err := unZip("./sample.zip", "./out")
	//err = unTar("./sample.tar", "./out")
	//if err != nil {
	//	log.Println("unTar return err", err)
	//}
	unarchive.Read("zip")
	sql.Open("mysql", "dbname")
}
