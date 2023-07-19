package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Journal struct {
	Title string
	Body string
}

func init() {
	fmt.Println("Server is running")
}

func main() {

	// mongodb related setup
	err := godotenv.Load()
    if err != nil {
    	log.Fatal(err)
    }

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("MONGODB_URI environment variable is not found")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// preparing and establishing server
	mux := http.NewServeMux()

	mux.HandleFunc("/", index)
	mux.HandleFunc("/test/data", provideData)
	mux.Handle("/html/styles/", http.StripPrefix("/html/styles/", http.FileServer(http.Dir("html/styles"))))

	err = http.ListenAndServe(":8000", mux)

	if err != nil {
		log.Fatal(err)
	}
}

func provideData(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)

	switch r.Method {
	case http.MethodGet:
		t, err := template.New("data.html").ParseFiles("./html/templates/data.html")

		if err != nil {
			log.Fatal(err)
		}

		err = t.Execute(w, nil)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)

	switch r.Method {
	case http.MethodGet:
		t, err := template.New("index.html").ParseFiles("./html/templates/index.html")

		if err != nil {
			log.Fatal(err)
		}

		data := Journal {
			Title: "Daily Dev Journal",
			Body: "Tesssst daemon watcher to my first taildwind work with my first respond using the template library again",
		}

		err = t.Execute(w, data)
		if err != nil {
			log.Fatal(err)
		}
	}
}