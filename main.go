package main

import (
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	http.HandleFunc("/home", handler)
	http.HandleFunc("/filter", filterHandler)
	http.HandleFunc("/paginate", paginate)
	http.ListenAndServe(":8080", nil)

}
