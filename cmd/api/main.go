package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rohansinghprogrammer/orders-api/internals/config"
	"github.com/rohansinghprogrammer/orders-api/internals/routes"
)

func main() {
	// Load Configurations
	cfg := config.MustLoadConfig()

	// Setup DB
	// Setup Router
	router := routes.OrderRoutes()

	// Setup HTTP Server
	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}
	// Setup Handlers

	// Start HTTP Server
	log.Printf("Server starting on http://localhost%s", cfg.Address)

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("Failed to start server %s", err.Error())
		}
	}()

	<- done

	// Handle Graceful Shutdown
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Failed to shut down server", slog.String("error", err.Error()))
	}
}
