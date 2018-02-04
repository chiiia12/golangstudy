package main

import (
	"fmt"
	"gopl.io/ch7/eval"
)

func main() {
	var formula string;
	//fmt.Scan(&formula)
	//fmt.Println(formula)
	formula = "pow(x, 3)"
	//TODO:変数の取得できん
	//r := regexp.MustCompile(`[A-Za-z]`)
	//fmt.Println(r.FindAllStringSubmatch(formula, -1))
	var x float64
	fmt.Println("input value for x")
	fmt.Scan(&x)
	env := eval.Env{"x": x}
	exp, err := eval.Parse(formula)
	if err != nil {
		fmt.Printf("Parse() error is not nil.error is: %v\n", err)
	}
	fmt.Printf("exp is %v\n", exp)
	fmt.Println(exp.Eval(env))

}
