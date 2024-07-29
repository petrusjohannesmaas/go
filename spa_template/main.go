package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "components/index.html")
	})

	mux.HandleFunc("/page_a", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "components/page_a.html")
	})

	mux.HandleFunc("/page_b", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "components/page_b.html")
	})

	fmt.Println("Server is running on port 80")
	http.ListenAndServe("localhost:8080", mux)
}
