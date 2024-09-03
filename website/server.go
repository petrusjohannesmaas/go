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

	mux.HandleFunc("/bio", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "frontend/bio.png")
	})

	mux.HandleFunc("/meta", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "frontend/meta.png")
	})

	mux.HandleFunc("/buffalo", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "frontend/buffalo.png")
	})

	mux.HandleFunc("/hopkins", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "frontend/hopkins.png")
	})

	mux.HandleFunc("/sydney", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "frontend/sydney.png")
	})

	fmt.Println("Server is running on port 80")
	http.ListenAndServe("localhost:8080", mux) // take out the localhost if you want to deploy with docker
}
