package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"github.com/madn-optimo/backend/handlers"
)

func SetupRoutes(
	router *chi.Mux,
	saleHandler *handlers.SaleHandler,
	purchaseOrderHandler *handlers.PurchaseOrderHandler,
) {
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
	}))

	router.Route("/api", func(r chi.Router) {
		r.Route("/sales", func(r chi.Router) {
			r.Get("/", saleHandler.GetAll)
			r.Get("/{id}", saleHandler.GetByID)
		})

		r.Route("/purchase-orders", func(r chi.Router) {
			r.Get("/", purchaseOrderHandler.GetAll)
			r.Get("/{id}", purchaseOrderHandler.GetByID)
			r.Post("/", purchaseOrderHandler.Create)
			r.Delete("/{id}", purchaseOrderHandler.Delete)
		})
	})
} 