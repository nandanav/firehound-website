package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func journalViewHandler(rw http.ResponseWriter, r *http.Request) {
	journalEntry := r.URL.Path[len("/journal/"):]
	fmt.Fprintf(rw, "<h1>Hello, welcome to journal entry %s<h1>", journalEntry)
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("static/css"))))

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		http.ServeFile(rw, r, "static/index.html")
	})

	http.HandleFunc("/favicon.svg", func(rw http.ResponseWriter, r *http.Request) {
		http.ServeFile(rw, r, "static/favicon.svg")
	})
	http.HandleFunc("/journal/", journalViewHandler)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
