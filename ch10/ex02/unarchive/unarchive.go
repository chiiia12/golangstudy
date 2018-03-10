package unarchive

import (
	"log"
	"fmt"
	"../zip"
	"../tar"
)

const OUTPUT_DIR = "./out"

type UnArchiver interface {
	UnArchive()
}

//file種別を持っておく配列を持っておく
var driver []string

func OpenUnArchiver(filetype string, inputfile string) (UnArchiver, error) {
	//配列から取り出して処理する
	//なかったらエラー吐く
	isExist := false
	for _, v := range driver {
		if v == inputfile {
			isExist = true
		}
	}
	if !isExist {
		return nil, fmt.Errorf("driver doesn't have %v", filetype)
	}
	switch(filetype) {
	case "zip":
		return &zip.ZipUnArchiver{inputfile, OUTPUT_DIR}, nil
	case "tar":
		return &tar.TarUnArchiver{inputfile, OUTPUT_DIR}, nil
	default:
		return nil, fmt.Errorf("driver doesn't have %v", filetype)
	}
}

func Register(name string) {
	//配列に登録
	driver = append(driver, name)
	log.Println(driver)
}
