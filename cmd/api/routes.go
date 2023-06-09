package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(app.cors)

	mux.Get("/login", app.authenticate)
	mux.Get("/", app.Home)
	mux.Get("/listings", app.Listings)

	return mux
}
