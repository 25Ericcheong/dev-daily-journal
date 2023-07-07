package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func init() {
	fmt.Println("Server is running")
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", index)

	err := http.ListenAndServe(":8000", mux)

	if err != nil {
		log.Fatal(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)

	switch r.Method {
	case http.MethodGet:
		t, err := template.New("home.html").ParseFiles("./templates/home.html")

		if err != nil {
			log.Fatal(err)
		}

		data := Article {
			Title: "Daily Dev Journal",
			Text: "Welcome to my first page with my first respond using the template library!",
		}

		err = t.Execute(w, data)
		if err != nil {
			log.Fatal(err)
		}
	}
}

type Article struct {
	Title string
	Text string
}