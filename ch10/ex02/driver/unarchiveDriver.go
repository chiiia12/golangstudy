package driver

import (
	"fmt"
	"../unarchive"
)


//file種別を持っておく配列を持っておく
var driver map[string]unarchive.UnArchiver

func init() {
	driver = make(map[string]unarchive.UnArchiver)
}

func OpenUnArchiver(filetype string, inputfile string) (unarchive.UnArchiver, error) {
	//配列から取り出して処理する
	//なかったらエラー吐く
	f, ok := driver[filetype]
	if !ok {
		return nil, fmt.Errorf("driver doesn't have %v", filetype)
	}
	return f, nil;
}

func Register(name string, u unarchive.UnArchiver) {
	driver[name] = u
}
