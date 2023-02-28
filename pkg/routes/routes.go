package routes

import (
	"myapp/handlers"
	"myapp/pkg/config"
	"net/http"
	"github.com/go-chi/chi"
)

func Routes(app *config.AppConfig) http.Handler {
	// Create multiplexer
	mux := chi.NewRouter()
	mux.Use(GenerateNoSurf)
	mux.Get("/", handlers.Repo.HomePage)
	mux.Get("/about", handlers.Repo.AboutPage)
	return mux
}
