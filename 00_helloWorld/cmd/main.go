package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Man-Crest/go--web-course/pkg/config"
	"github.com/Man-Crest/go--web-course/pkg/handlers"
	"github.com/Man-Crest/go--web-course/pkg/render"

	"github.com/alexedwards/scs/v2"
)

var PortNumber string = ":8080"
var session *scs.SessionManager

func main() {

	var app config.AppConfig

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = false
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

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
