package unarchive

import "fmt"

type UnArchive struct {
}

//file種別を持っておく配列を持っておく
var driver []string

func (file *UnArchive) Read(name string) {
	//配列から取り出して処理する
	//なかったらエラー吐く
	for _, v := range driver {
		if v == name {
			//何かする
			return
		}
	}
	fmt.Errorf("driver doesn't have %v", name)
}

func Register(name string) {
	//配列に登録
	driver = append(driver, name)
}
