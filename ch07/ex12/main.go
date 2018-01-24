// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 192.

// Http2 is an e-commerce server with /list and /price endpoints.
package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
)

var datalist = template.Must(template.New("dataList").Parse(`<table>
	<tr>
		<td>item</td>
		<td>price</td>
	</tr>
	{{range $key, $value := .}}
	<tr>
		<td>{{ $key }}</td>
		<td>{{ $value }}</td>
	</tr>
	{{end}}
</table>`))



func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//!-main

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	if err := datalist.Execute(w, db); err != nil {
		log.Fatal(err)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

