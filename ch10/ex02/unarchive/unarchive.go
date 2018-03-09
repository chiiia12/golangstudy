package unarchive

import (
	"fmt"
	"log"
)

type UnArchive struct {
}

//file種別を持っておく配列を持っておく
var driver []string

func Read(name string) {
	//配列から取り出して処理する
	//なかったらエラー吐く
	for _, v := range driver {
		if v == name {
			//何かする
			log.Println("見つかった")
			return
		}
	}
	fmt.Errorf("driver doesn't have %v", name)
}

func Register(name string) {
	//配列に登録
	driver = append(driver, name)
	log.Println(driver)
}
