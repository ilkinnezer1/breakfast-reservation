package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"myapp/pkg/config"
	"myapp/pkg/models"
	"net/http"
	"path/filepath"
)

// HandleError handles with errors
func HandleError(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}


var app *config.AppConfig
func CreateNewTemplates(a *config.AppConfig){
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData{
	return td
}

func RenderTemplate(w http.ResponseWriter, html string, td *models.TemplateData) {
	// TODO:
	//Create the template cache
	var tempCache map[string]*template.Template

	if app.UseCache {
		tempCache = app.TemplateCache
	} else{
		tempCache, _ = CreateTemplateCache()
	}
	// Get request template cache
	temp, ok := tempCache[html]
	if !ok {
		log.Fatal("Couldn't get the template cache")
	}
	// Hold bytes
	buffer := new(bytes.Buffer)

	td = AddDefaultData(td)

	err := temp.Execute(buffer, td)
	HandleError(err)
	// render the template
	_, err = buffer.WriteTo(w)
	HandleError(err)
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	// Create entire cache
	// Get all files from templates folder
	pages, err := filepath.Glob("./templates/*.html")
	HandleError(err)
	// Range thru pages
	for _, page := range pages {
		name := filepath.Base(page)
		tempSet, err := template.New(name).ParseFiles(page)
		HandleError(err)

		matches, err := filepath.Glob("./templates/*.layout.html")
		HandleError(err)

		if len(matches) > 0 {
			tempSet, err = tempSet.ParseGlob("./templates/*.layout.html")
		}

		cache[name] = tempSet
	}

	return cache, nil

}
