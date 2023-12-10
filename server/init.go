package server

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func StartHTTPServer() error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Get("/api/s3/upload", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("upload"))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Println("Starting Relay server on :" + port)

	if err := http.ListenAndServe(":"+port, r); err != nil {
		return err
	}

	return nil
}
