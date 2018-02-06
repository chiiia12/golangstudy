package main

import (
	"net/http"
	"log"
	"fmt"
	"gopl.io/ch7/eval"
)

func main() {
	fmt.Println("ex) http://localhost:8000?expr=(1.1 * 2 %%2b 3)*3&expr=355/113")
	http.HandleFunc("/", calc)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
func calc(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	values, ok := r.Form["expr"]
	if !ok {

	}
	fmt.Println(values)

	for _, v := range values {
		env := eval.Env{}
		exp, err := eval.Parse(v)
		if err != nil {
			fmt.Printf("Parse error is not nil.error is :%v\n", err)
		}
		fmt.Fprintf(w, "answer is %v\n", exp.Eval(env))
	}

}
