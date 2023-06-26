package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server is running")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Server has responded. You have requested %s\n", r.URL.Path)
	})

	http.ListenAndServe(":8000", nil)
}
