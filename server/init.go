package server

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	opt "github.com/gursheyss/relay/options" // Aliased import
)

func SetupServer(opts ...opt.Option) (func() error, error) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	for _, option := range opts {
		switch option {
		case opt.S3:
			// Define your routes for S3 here
		case opt.Lambda:
			// Define your routes for Lambda here
		default:
			// Handle unknown options if needed
		}
	}

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
