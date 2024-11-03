package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	// Use the CORS middleware to specify who is allowed to connect.
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Use the Heartbeat middleware to add a "/ping" route for health checks.
	mux.Use(middleware.Heartbeat("/ping"))
	// status: 200OK
	// "."

	// Add a basic route for the root path "/"
	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Broker service!"))
	})

	mux.Post("/", app.Broker)

	mux.Post("/handle", app.HandleSubmission)

	return mux
}
