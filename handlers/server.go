package handlers

import (
	"fmt"
	"log"
	"myapp/pkg/config"
	"myapp/render"
	"net/http"
)
const port = ":9234"

func Start(){
	var app config.AppConfig

	templCache, err := render.CreateTemplateCache()
	if err != nil { 
		log.Fatal(err)
	}
	app.TemplateCache = templCache
	app.UseCache = false
	// Use repo pattern
	repo := CreateNewRepo(&app)
	CreateNewHandlers(repo)


	render.CreateNewTemplates(&app)
	http.HandleFunc("/", Repo.HomePage)
	http.HandleFunc("/about", Repo.AboutPage)
	fmt.Println(fmt.Sprintf("Server is running on port %s", port))
	http.ListenAndServe(port, nil)
}