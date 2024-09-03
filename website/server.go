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

	mux.HandleFunc("/react", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "frontend/react.png")
	})

	mux.HandleFunc("/go", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "frontend/go.png")
	})

	mux.HandleFunc("/solidity", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "frontend/solidity.png")
	})

	mux.HandleFunc("/docker", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "frontend/docker.png")
	})

	fmt.Println("Server is running on port 80")
	http.ListenAndServe("localhost:8080", mux) // take out the localhost if you want to deploy with docker
}
