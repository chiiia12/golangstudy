package ex16

import (
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/", calc)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
func calc(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	queryParam := r.URL.Query()

}
