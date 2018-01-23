package main

import (
	"net/http"
	"fmt"
	"log"
	"strconv"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}

type dollars float32

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "url is ", req.URL)
	switch req.URL.Path {
	case "/create":
		item := req.URL.Query().Get("item")
		fmt.Fprintf(w, "price is %v", req.URL.Query().Get("price"))
		price, _ := strconv.ParseFloat(req.URL.Query().Get("price"), 32)
		db[item] = dollars(price)
		fmt.Fprintf(w, "create success item:%v price:%v", item, price)
	case "/read":
		for item, price := range db {
			fmt.Fprintf(w, "%s:%s\n", item, price)
		}
	case "/update":
		item := req.URL.Query().Get("item")
		fmt.Fprintf(w, "price is %v\n", req.URL.Query().Get("price"))
		price, _ := strconv.ParseFloat(req.URL.Query().Get("price"), 32)
		if _, ok := db[item]; !ok {
			fmt.Fprintf(w, "ERROR:db doesn't have this item:%v", item)
		}
		db[item] = dollars(price)
		fmt.Fprintf(w, "update success item:%v price:%v", item, price)
	case "/delete":
		item := req.URL.Query().Get("item")
		fmt.Fprintf(w, "price is %v", req.URL.Query().Get("price"))
		price, _ := strconv.ParseFloat(req.URL.Query().Get("price"), 32)
		if _, ok := db[item]; !ok {
			fmt.Fprintf(w, "ERROR:db doesn't have this item:%v", item)
		}
		delete(db, item)
		fmt.Fprintf(w, "delete success item:%v price:%v", item, price)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such page:%s\n", req.URL)

	}
}
