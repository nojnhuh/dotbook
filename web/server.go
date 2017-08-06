// Package web conatins the web server
package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/nojnhuh/dotbook/db"
)

var (
	tmplDotbook     *template.Template
	tmplDotbookList *template.Template
)

// handler simply retrieves a default dotbook and displays it.
func dotbookHandler(w http.ResponseWriter, r *http.Request) {
	dbName := r.URL.Query()["q"][0]
	book := db.GetDotbook(dbName)
	tmplDotbook.ExecuteTemplate(w, "base", book)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	books := db.GetAllDotbooks()
	tmplDotbookList.ExecuteTemplate(w, "base", books)
}

func setPaths() {
	http.HandleFunc("/dotbook", dotbookHandler)
	http.HandleFunc("/", indexHandler)
}

func parseTemplates() {
	var err error
	tmplDotbook, err = template.New("web/static/templates/dotbook.gohtml").Funcs(template.FuncMap{
		"half": func(n float64) float64 { return n / 2 },
	}).ParseFiles("web/static/templates/base.gohtml", "web/static/templates/dotbook.gohtml")
	if err != nil {
		panic(err)
	}
	tmplDotbookList, err = template.ParseFiles("web/static/templates/base.gohtml", "web/static/templates/dotbook_list.gohtml")
	if err != nil {
		panic(err)
	}
}

// InitServer starts the web serve and declares handler functions
func InitServer(port int) {
	setPaths()
	parseTemplates()
	log.Printf("Ready to serve HTTP on port %d.", port)
	path := fmt.Sprintf(":%d", port)
	log.Fatal(http.ListenAndServe(path, nil))
}
