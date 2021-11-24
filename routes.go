package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *application) routes() *chi.Mux {
	// add middleware

	// add routes
	a.get("/", a.Handlers.Home)

	// static assets
	fileServer := http.FileServer(http.Dir("./public"))
	// add these routes to the celeritas routes
	a.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileServer))

	return a.App.Routes
}
