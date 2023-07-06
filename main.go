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

	const tpl = `
	<!DOCTYPE html>
	<html>
		<head>
			<meta charset="UTF-8">
			<title>{{.Title}}</title>
		</head>
		<body>
			{{.Text}}
		</body>
	</html>
	`

	switch r.Method {
	case http.MethodGet:
		t, err := template.New("home").Parse(tpl)

		if err != nil {
			log.Fatal(err)
		}

		data := struct {
			Title string
			Text  string
		}{
			Title: "Daily Dev Journal",
			Text:  "Welcome to my first page with my first respond using the template library!",
		}

		err = t.Execute(w, data)
		if err != nil {
			log.Fatal(err)
		}
	}
}
