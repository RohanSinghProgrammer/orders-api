package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	orderhandlers "github.com/rohansinghprogrammer/orders-api/internals/handlers/order-handlers"
)

func OrderRoutes() *chi.Mux {
	router := chi.NewMux()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/orders", orderHandlers)

	return router
}

func orderHandlers(router chi.Router) {
	handlers := &orderhandlers.Order{}

	router.Get("/", handlers.List)
	router.Get("/{id}", handlers.GetByID)
	router.Post("/", handlers.Create)
	router.Put("/{id}", handlers.UpdateByID)
	router.Delete("/{id}", handlers.DeleteByID)
}