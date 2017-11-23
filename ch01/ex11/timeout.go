package main

import (
	"fmt"
	"net/http"
	"time"
	"log"
)

func main() {
	fmt.Println("timeout.go")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1000 * time.Second)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
