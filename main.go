package main

import (
	"html/template"
	"net/http"

	"google.golang.org/appengine"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/link", linkHandler)
	http.HandleFunc("/img", imgHandler)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	appengine.Main()
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index", nil)
}

func linkHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "link", nil)
}

func imgHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "img", nil)
}

func renderTemplate(w http.ResponseWriter, tmpl string, pageVars interface{}) {
	err := templates.ExecuteTemplate(w, tmpl+".html", pageVars)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
