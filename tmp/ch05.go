package main

import (
	"os"
	"fmt"
)

func main() {
	//ループ変数の補足
	var rmdirs []func()
	dirs := []string{"hoge", "piyo"}
	for i := 0; i < len(dirs); i++ {
		os.MkdirAll(dirs[i], 0755)
		rmdirs = append(rmdirs, func() {
			//定義しているだけで実行はしていない
			os.RemoveAll(dirs[i])
		})
	}

	//recover nest
	recoverSample()

	//ポイントレシーバー
	//変数に入れずにメソッドを呼ぼうとするとアドレス化可能じゃないからコンパイルエラー
	p := Point{1, 2}
	p.pointMethod()



}


type Point struct {
	a, b int
}

func (p *Point) pointMethod() {

}

func recoverSample() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Errorf("internal error:%v ", p)
		}
	}()
	recoverSample2()
}
func recoverSample2() {
	panic("panic")

}
