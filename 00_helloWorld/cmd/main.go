package main

import (
	"fmt"
	"net/http"

	"github.com/Man-Crest/go--web-course/pkg/config"
	"github.com/Man-Crest/go--web-course/pkg/handlers"
	"github.com/Man-Crest/go--web-course/pkg/render"
)

var PortNumber string = ":8080"

func main() {

	var app config.AppConfig

	ts, err := render.CreateTempleteCache()
	if err != nil {
		fmt.Println(err)
	}

	app.TemplateCache = ts

	render.NewTemplates(&app)
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("starting application on port number %s", PortNumber))
	_ = http.ListenAndServe(PortNumber, nil)
}
