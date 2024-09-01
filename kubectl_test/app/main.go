package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello human! \n")
		fmt.Println("A human has been greeted ✓")
	})

	mux.HandleFunc("/joke", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "your mom \n")
		fmt.Println("A joke has been made ✓")
	})

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", mux)
}
