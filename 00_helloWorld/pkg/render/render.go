package render

import (
	"bytes"
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/Man-Crest/go--web-course/pkg/config"
	"github.com/Man-Crest/go--web-course/pkg/models"
)

// func RenderTemplate(w http.ResponseWriter, tmpl string) {
// 	parsedTmpl, err := template.ParseFiles(fmt.Sprintf("../templates/%s", tmpl), "../templates/base.layout.html")
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Error parsing template: %s", err), http.StatusInternalServerError)
// 		return
// 	}

// 	err = parsedTmpl.Execute(w, nil)
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Error executing template: %s", err), http.StatusInternalServerError)
// 		return
// 	}
// }

// var tc = make(map[string]*template.Template)

// func RenderTemplate(w http.ResponseWriter, t string) {
// 	var tmpl *template.Template
// 	var err error
// 	_, inMap := tc[t]

// 	if !inMap {
// 		fmt.Println("Creating new cache")
// 		err = CreateParsedTemplateCache(w, t)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 	} else {
// 		fmt.Println("already created and using stored template")
// 	}
// 	tmpl = tc[t]

// 	err = tmpl.Execute(w, nil)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	var ts map[string]*template.Template
	if app.UseCache {
		ts = app.TemplateCache
	} else {
		ts, _ = CreateTempleteCache()
	}
	app.UseCache = true

	t, ok := ts[tmpl]
	var err error
	if !ok {
		fmt.Println(err)
	}

	buf := new(bytes.Buffer)

	err = t.Execute(buf, td)
	if err != nil {
		fmt.Println(err)
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println(err)
	}
}

// func CreateParsedTemplateCache(w http.ResponseWriter, t string) error {
// 	templates := []string{
// 		fmt.Sprintf("../templates/%s", t),
// 		"../templates/base.layout.html",
// 	}
// 	parsedTmpl, err := template.ParseFiles(templates...)

// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	err = parsedTmpl.Execute(w, nil)

// 	tc[t] = parsedTmpl
// 	return err
// }

func CreateTempleteCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("../templates/*.page.templ")
	fmt.Println(pages)
	if err != nil {
		fmt.Println(err)
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("../templates/*.layout.templ")

		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("../templates/*layout.templ")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}
	return myCache, nil

}
