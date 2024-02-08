package handlers

import (
	"net/http"

	"github.com/Man-Crest/go--web-course/pkg/config"
	"github.com/Man-Crest/go--web-course/pkg/models"
	"github.com/Man-Crest/go--web-course/pkg/render"
)

var Repo *Reposatory

type Reposatory struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Reposatory {
	return &Reposatory{
		App: a,
	}
}

func NewHandlers(r *Reposatory) {
	Repo = r
}

func (m *Reposatory) Home(w http.ResponseWriter, r *http.Request) {

	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)

	stringMap1 := make(map[string]string)
	stringMap1["test"] = "Hello, home"
	render.RenderTemplate(w, "home.page.templ", &models.TemplateData{
		StringMap: stringMap1,
	})
}

func (m *Reposatory) About(w http.ResponseWriter, r *http.Request) {
	stringMap2 := make(map[string]string)
	stringMap2["test"] = "Hello, again"

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap2["remote_ip"] = remoteIp
	render.RenderTemplate(w, "about.page.templ", &models.TemplateData{
		StringMap: stringMap2,
	})
}
