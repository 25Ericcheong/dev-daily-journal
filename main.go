package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"

	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Journal struct {
	Title string
	Summary string
	Body string
}

const noPanic = false
const mongoEnvUri = "MONGODB_URI"
const port = ":8000"
const localhost = "http://localhost:8000/"

func init() {
	fmt.Println("Server is running")
}

func main() {
	fmt.Printf("Running on port %v", port)

	err := godotenv.Load()
    if err != nil {
		log.Fatal("Failed to load relevant environment variables", err)
	}

	uri := os.Getenv(mongoEnvUri)
	if uri == "" {
		log.Fatal("MONGODB_URI environment variable is not found")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Failed to connect to MongoDB Client", err)
	}


	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// preparing and establishing server
	mux := http.NewServeMux()

	mux.HandleFunc("/", index)
	mux.HandleFunc("/journal/", journal)

	// enable server to serve output of tailwind css when requested and htmx
	mux.Handle("/html/styles/", http.StripPrefix("/html/styles/", http.FileServer(http.Dir("./html/styles"))))
	mux.Handle("/html/htmx/", http.StripPrefix("/html/htmx/", http.FileServer(http.Dir("./html/htmx"))))

	// auto open default browser if linux
	if runtime.GOOS == "linux" {
		go exec.Command("xdg-open", localhost).Start()
	}

	// start server and listen to provided port
	err = http.ListenAndServe(port, mux)
	if (err != nil) {
		log.Fatal(err)
	}
	
}

func journal(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	fmt.Println(r.Method)

	switch r.Method {
	case http.MethodGet:
		temp, err := template.ParseFiles("./html/templates/journals/journal.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = temp.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	fmt.Println(r.Method)

	switch r.Method {
	case http.MethodGet:
		
		temp, err := template.ParseFiles(
			"./html/templates/index.html",
			"./html/templates/parts/head.html",
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := Journal {
			Title: "Daily Dev Journal",
			Body: "Temporary dummy data for not - will include data from MongoDB",
		}

		err = temp.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}