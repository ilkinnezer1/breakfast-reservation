package handlers

import (
	"net/http"
	"myapp/render"
)
// HomePage is the home page handler
func HomePage(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, "home.html")
}
func AboutPage(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, "about.html")
}

