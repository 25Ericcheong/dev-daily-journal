package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Server is running")

	dir := http.Dir("./static")
	fs := http.FileServer(dir)
	mux := http.NewServeMux()
	mux.Handle("/test", fs)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Server has responded. You have requested %s\n", r.URL.Path)
	})

	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		log.Fatal(err)
	}
}
