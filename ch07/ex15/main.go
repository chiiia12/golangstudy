package main

import (
	"fmt"
	"gopl.io/ch7/eval"
)

func main() {
	var formula string;
	for {
		fmt.Printf("input a expression:")
		fmt.Scan(&formula)
		fmt.Println(formula)

		env := eval.Env{}
		exp, err := eval.Parse(formula)
		if err != nil {
			fmt.Printf("Parse() error is not nil.error is: %v\n", err)
		}
		fmt.Println(exp.Eval(env))
	}
}

//変数
