package main

import (
	"os"
	"fmt"
	"log"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World! This is a test handler")
}

func journalViewHandler(w http.ResponseWriter, r *http.Request) {
	journalEntry := r.URL.Path[len("/journal/"):]
	fmt.Fprintf(w, "<h1>Hello, welcome to journal entry %s<h1>", journalEntry)
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/journal/", journalViewHandler)

	log.Fatal(http.ListenAndServe(":" + port, nil))
}