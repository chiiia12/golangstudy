package main

//コンパイラは通るけど呼ぶと落ちるパターン
func max(vals ...int) int {
	if len(vals) == 0 {
		panic("param size is 0")
	}
	max := vals[0]
	for _, val := range vals {
		if max < val {
			max = val
		}
	}
	return max
}

//コンパイラで引数なしパターンのエラーがわかる
//func max(x int,vals...int)int
//minやmaxの初期値を0にしてると不十分
//配列の0番目にするか、intの一番マイナス値をとるなど？
func min(vals ...int) int {
	if len(vals) == 0 {
		panic("param size is 0")
	}
	min := vals[0]
	for _, val := range vals {
		if min > val {
			min = val
		}
	}
	return min
}
