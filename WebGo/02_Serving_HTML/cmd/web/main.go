package main

import (
	"fmt"
	"log"
	"net/http"
	"templates/pkg/config"
	"templates/pkg/handlers"
	"templates/pkg/render"
)
const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCahce = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)

	render.NewTemplates(&app)

	/*
	http.HandleFunc("/", handlers.Repo.HomePage)
	http.HandleFunc("/about", handlers.Repo.AboutPage)
	*/
	
	fmt.Println("Server running on port", portNumber)
	/*
	err = http.ListenAndServe(portNumber, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
	*/

	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}