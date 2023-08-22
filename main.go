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

const noPanic = false
const mongoEnvUri = "MONGODB_URI"
const port = ":8000"

func init() {
	fmt.Println("Server is running")
}

func checkErr(err error, toPanic bool) {
	if err != nil {
		if toPanic == true {
			panic(err)
		} else {
			log.Fatal(err)
		}
    }
}

func main() {

	// mongodb related setup
	err := godotenv.Load()
    checkErr(err, noPanic)

	uri := os.Getenv(mongoEnvUri)
	if uri == "" {
		log.Fatal("MONGODB_URI environment variable is not found")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	checkErr(err, !noPanic)

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// preparing and establishing server
	mux := http.NewServeMux()

	mux.HandleFunc("/", index)
	mux.HandleFunc("/journal/create", createJournal)
	mux.Handle("/html/styles/", http.StripPrefix("/html/styles/", http.FileServer(http.Dir("./html/styles"))))

	err = http.ListenAndServe(port, mux)
	checkErr(err, noPanic)
}

func createJournal(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	fmt.Println(r.Method)

	switch r.Method {
	case http.MethodGet:
		temp, err := template.New("journal.html").ParseFiles("./html/templates/journal.html")
		checkErr(err, noPanic)

		err = temp.Execute(w, nil)
		checkErr(err, noPanic)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	fmt.Println(r.Method)

	switch r.Method {
	case http.MethodGet:
		temp, err := template.New("index.html").ParseFiles("./html/templates/index.html")
		checkErr(err, noPanic)

		data := Journal {
			Title: "Daily Dev Journal",
			Body: "Temporary dummy data for not - will include data from MongoDB",
		}

		err = temp.Execute(w, data)
		checkErr(err, noPanic)

	case http.MethodPost:
		err := r.ParseForm()
		checkErr(err, noPanic)

		newJournal := Journal{}
		newJournal.Title = r.Form.Get("new_journal_title")
		newJournal.Body = r.Form.Get("new_journal_body")

	}
}