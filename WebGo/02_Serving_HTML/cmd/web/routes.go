package main

import (
	"net/http"
	"templates/pkg/config"
	"templates/pkg/handlers"

	"github.com/go-chi/chi/v5"
)

func routes(app *config.AppConfig) http.Handler {
	
	/*
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.HomePage))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.AboutPage))
	*/

	mux := chi.NewRouter()
	mux.Get("/", handlers.Repo.HomePage)
	mux.Get("/about", handlers.Repo.AboutPage)

	return mux
}

