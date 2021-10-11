package main

import (
	// "fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type UpdatesPage struct {
	IsJournal bool
	IsGoals   bool
	Data      interface{}
}

type Journal struct {
	Title string
}

var tmpls = make(map[string]*template.Template)

func staticFileViewHandler(rw http.ResponseWriter, r *http.Request, file string, data interface{}) {
	err := tmpls[file].ExecuteTemplate(rw, file, data)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func journalViewHandler(rw http.ResponseWriter, r *http.Request) {
	// journalEntry := r.URL.Path[len("/journal/"):]

	journal := Journal{Title: "Testing"}

	t, _ := template.ParseFiles("templates/updates.tmpl.html")
	err := t.Execute(rw, journal)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	port := os.Getenv("PORT")
	tmpls["index.html"] = template.Must(template.ParseFiles("templates/header.tmpl.html", "static/index.html"))
	tmpls["timeline.html"] = template.Must(template.ParseFiles("templates/header.tmpl.html", "static/timeline.html"))
	tmpls["advisors.html"] = template.Must(template.ParseFiles("templates/header.tmpl.html", "static/advisors.html"))
	tmpls["updates.html"] = template.Must(template.ParseFiles("templates/header.tmpl.html", "templates/journal.tmpl.html", "static/updates.html"))

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("static/css"))))

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		staticFileViewHandler(rw, r, "index.html", "testing")
	})

	http.HandleFunc("/timeline.html", func(rw http.ResponseWriter, r *http.Request) {
		staticFileViewHandler(rw, r, "timeline.html", "testing")
	})

	http.HandleFunc("/advisors.html", func(rw http.ResponseWriter, r *http.Request) {
		staticFileViewHandler(rw, r, "advisors.html", "testing")
	})

	http.HandleFunc("/updates.html", func(rw http.ResponseWriter, r *http.Request) {
		updatePageConfig := UpdatesPage{IsJournal: true, IsGoals: false, Data: Journal{Title: "testing"}}
		staticFileViewHandler(rw, r, "updates.html", updatePageConfig)
	})

	http.HandleFunc("/favicon.svg", func(rw http.ResponseWriter, r *http.Request) {
		http.ServeFile(rw, r, "static/favicon.svg")
	})

	http.HandleFunc("/journal/", journalViewHandler)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
