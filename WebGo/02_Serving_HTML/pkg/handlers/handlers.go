package handlers

import (
	"net/http"
	"templates/pkg/config"
	"templates/pkg/models"
	"templates/pkg/render"
)

// Repo is the repository used by the handler
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository {
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandler(r * Repository) {
	Repo = r
}

func (m *Repository) HomePage (w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}
func (m *Repository) AboutPage (w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	// send the data to the template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}