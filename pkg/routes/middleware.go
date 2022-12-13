package routes

import (
	"myapp/pkg/config"
	"net/http"

	"github.com/justinas/nosurf"
)

var app config.AppConfig

func GenerateNoSurf(next http.Handler) http.Handler{
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		Secure: app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

func LoadSession(next http.Handler) http.Handler{
	return app.Session.LoadAndSave(next)
}