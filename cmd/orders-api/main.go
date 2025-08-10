package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rohansinghprogrammer/orders-api/internals/config"
)

func main() {
	// Load Configurations
	cfg := config.MustLoadConfig()
	// Setup DB
	// Setup Router
	router := chi.NewRouter()
	// Initialize Logger
	router.Use(middleware.Logger)
	// Setup HTTP Server
	server := http.Server{
		Addr: cfg.Address,
		Handler: router,
	}
	// Setup Handlers
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Go"))
	})
	// Start HTTP Server
	log.Printf("Server starting on http://localhost%s",cfg.Address)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server %s", err.Error())
	}
	// Handle Graceful Shutdown
}