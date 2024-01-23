package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var PortNumber string = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.html")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTmpl, err := template.ParseFiles(fmt.Sprintf("./templates/%s", tmpl))
	if err != nil {
		log.Fatal(err)
	}
	err = parsedTmpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("starting application on port number %s", PortNumber))
	_ = http.ListenAndServe(PortNumber, nil)
}
