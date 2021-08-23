package handler

import (
	"log"
	"net/http"
	"text/template"
)

// Home page handler : http://localhost:5000/
func homeHandler(w http.ResponseWriter, r *http.Request) {
	webTemplate, _ := template.ParseFiles("./templates/home.html")
	title := "Home page"
	webTemplate.Execute(w, title)
}

// About page handler : http://localhost:5000/about
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	webTemplate, _ := template.ParseFiles("./templates/about.html")
	title := "About page"
	webTemplate.Execute(w, title)
}

func RequestHandler() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about/", aboutHandler)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
