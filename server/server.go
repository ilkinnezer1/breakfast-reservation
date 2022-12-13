package server

import (
	"fmt"
	"log"
	"myapp/handlers"
	"myapp/pkg/config"
	"myapp/pkg/routes"
	"myapp/render"
	"net/http"
	"time"
	"github.com/alexedwards/scs/v2"
)

const port = ":9234"
var session *scs.SessionManager

func Start() {
	var app config.AppConfig
    // Session configuration
	// TODO when in production changes to true
	app.InProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	templCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	app.TemplateCache = templCache
	app.UseCache = false
	// Use repo pattern
	repo := handlers.CreateNewRepo(&app)
	handlers.CreateNewHandlers(repo)
	render.CreateNewTemplates(&app)

	server := &http.Server{
		Addr:    port,
		Handler: routes.Routes(&app),
	}

	err = server.ListenAndServe()
	fmt.Println(fmt.Sprintf("Server is running on port %s", port))
	log.Fatal(err)
}
