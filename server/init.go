package server

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupServer() (func() error, error) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	startServer := func() error {
		log.Println("Starting Relay server on port " + port)
		return http.ListenAndServe(":"+port, r)
	}

	return startServer, nil
}
