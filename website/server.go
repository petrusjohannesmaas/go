package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "frontend/portfolio.html")
	})

	mux.HandleFunc("/style", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "frontend/style.css")
	})

	fmt.Println("Server is running on port 80")
	http.ListenAndServe("localhost:8080", mux) // take out the localhost if you want to deploy with docker
}
