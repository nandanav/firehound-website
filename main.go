package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/carlmjohnson/gateway"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World! This is a test handler")
}

func journalViewHandler(w http.ResponseWriter, r *http.Request) {
	journalEntry := r.URL.Path[len("/journal/"):]
	fmt.Fprintf(w, "<h1>Hello, welcome to journal entry %s<h1>", journalEntry)
}

func main() {
	port := flag.Int("port", -1, "specify a port to use http rather than AWS Lambda")
	flag.Parse()
	listener := gateway.ListenAndServe
	portStr := "n/a"
	if *port != -1 {
		portStr = fmt.Sprintf(":%d", *port)
		listener = http.ListenAndServe
		// http.Handle("/", http.FileServer(http.Dir("./public")))
	}

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/journal/", journalViewHandler)

	log.Fatal(listener(portStr, nil))
}