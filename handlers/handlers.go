package handlers

import (
	"myapp/pkg/config"
	"myapp/pkg/models"
	"myapp/render"
	"net/http"
)

// Use Repository Pattern

// Used by handlers
var Repo *Repository

type Repository struct {
	App *config.AppConfig
}
// CreateNewRepo creates new repository
func CreateNewRepo(a *config.AppConfig) *Repository { 
	return &Repository{
		App: a,
	}
}
 //CreateNewHandlers sets the repo for the handlers
func CreateNewHandlers(r *Repository){
	Repo = r
}

// HomePage is the home page handler
func (m *Repository) HomePage(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, "home.html", &models.TemplateData{})
}
func (m *Repository) AboutPage(w http.ResponseWriter, r *http.Request){
	stringMap := make(map[string]string)
	stringMap["reserved"] = "Welcome to the Reservation Centerr"

	render.RenderTemplate(w, "about.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

