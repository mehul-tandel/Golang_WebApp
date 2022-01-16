package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mehul-tandel/Golang_WebApp/pkg/config"
	"github.com/mehul-tandel/Golang_WebApp/pkg/handlers"
	"github.com/mehul-tandel/Golang_WebApp/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false // development mode creates template cache everytime

	// create and set Repo in handlers
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// Pass current config instance to app var in render
	render.NewTemplates(&app)

	// // assign handler functions to routes (before using pat)
	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting application on port %s", portNumber)
	// http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

}
