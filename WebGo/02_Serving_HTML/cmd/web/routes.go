package main

import (
	"net/http"
	"templates/pkg/config"
	"templates/pkg/handlers"

	"github.com/bmizerany/pat"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.HomePage))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.AboutPage))

	return mux
}

