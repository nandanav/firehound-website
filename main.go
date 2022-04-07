package main

import (
	// "fmt"
	"fmt"
	"html/template"
	"io/ioutil"
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
	JournalNumber int
	Title         string
	PdfFilePath   string
}

var tmpls = make(map[string]*template.Template)

func staticFileViewHandler(rw http.ResponseWriter, r *http.Request, file string, data interface{}) {
	err := tmpls[file].ExecuteTemplate(rw, file, data)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	port := os.Getenv("PORT")
	tmpls["index.html"] = template.Must(template.ParseFiles("templates/header.tmpl.html", "static/index.html"))
	tmpls["timeline.html"] = template.Must(template.ParseFiles("templates/header.tmpl.html", "static/timeline.html"))
	tmpls["advisors.html"] = template.Must(template.ParseFiles("templates/header.tmpl.html", "static/advisors.html"))
	tmpls["updates.html"] = template.Must(template.ParseFiles("templates/header.tmpl.html", "templates/journals.tmpl.html", "templates/goals.tmpl.html", "static/updates.html"))

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("static/css"))))
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("static/assets"))))

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
		files, fileReadErr := ioutil.ReadDir("static/assets/journals/")

		if fileReadErr != nil {
			log.Fatal(fileReadErr, ": failed to read PDF file directory")
		}
		journals := make([]Journal, len(files))
		for fileNumber, file := range files {
			journals[fileNumber] = Journal{
				JournalNumber: fileNumber + 1,
				Title:         file.Name()[2:],
				PdfFilePath:   fmt.Sprintf("/assets/journals/%s", file.Name()),
			}
		}

		updatePageConfig := UpdatesPage{IsJournal: true, IsGoals: false, Data: journals}
		staticFileViewHandler(rw, r, "updates.html", updatePageConfig)
	})

	http.HandleFunc("/updates/journals", func(rw http.ResponseWriter, r *http.Request) {
		files, fileReadErr := ioutil.ReadDir("static/assets/journals/")

		if fileReadErr != nil {
			log.Fatal(fileReadErr, ": failed to read PDF file directory")
		}
		journals := make([]Journal, len(files))
		for fileNumber, file := range files {
			journals[fileNumber] = Journal{
				JournalNumber: fileNumber + 1,
				Title:         file.Name()[2:],
				PdfFilePath:   fmt.Sprintf("/assets/journals/%s", file.Name()),
			}
		}

		fmt.Println(journals)

		updatePageConfig := UpdatesPage{IsJournal: true, IsGoals: false, Data: journals}
		staticFileViewHandler(rw, r, "updates.html", updatePageConfig)
	})

	http.HandleFunc("/updates/goals", func(rw http.ResponseWriter, r *http.Request) {
		updatePageConfig := UpdatesPage{IsJournal: false, IsGoals: true, Data: Journal{Title: "testing goals"}}
		staticFileViewHandler(rw, r, "updates.html", updatePageConfig)
	})

	http.HandleFunc("/updates/bibliography.html", func(rw http.ResponseWriter, r *http.Request) {
		updatePageConfig := UpdatesPage{IsJournal: false, IsGoals: false, Data: nil}
		staticFileViewHandler(rw, r, "updates.html", updatePageConfig)
	})

	http.HandleFunc("/favicon.svg", func(rw http.ResponseWriter, r *http.Request) {
		http.ServeFile(rw, r, "static/favicon.svg")
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
