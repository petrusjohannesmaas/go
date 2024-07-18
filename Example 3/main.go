package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	fmt.Println("Server started on port 8080")
	fmt.Println("Run 'curl http://localhost:8080/books/HarryPotter/page/1' in cmd to test endpoint")
	http.ListenAndServe("localhost:8080", r)
}
