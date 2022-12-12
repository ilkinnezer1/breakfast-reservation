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

	render.CreateNewTemplates(&app)
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/about", AboutPage)
	fmt.Println(fmt.Sprintf("Server is running on port %s", port))
	http.ListenAndServe(port, nil)
}