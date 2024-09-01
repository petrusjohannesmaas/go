package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello human!")
		fmt.Println("A human has been greeted ✓")
	})

	mux.HandleFunc("/joke", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "your mom")
		fmt.Println("A joke has been made ✓")
	})

	fmt.Println("Server is running on port 80")
	http.ListenAndServe("localhost:8080", mux)
}
