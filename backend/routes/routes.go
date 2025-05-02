package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/madn-optimo/backend/handlers"
)

func SetupRoutes(
	router *chi.Mux,
	saleHandler *handlers.SaleHandler,
	purchaseOrderHandler *handlers.PurchaseOrderHandler,
) {
	// Middleware
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)

	// Enable CORS
	router.Use(middleware.AllowContentType("application/json"))
	router.Use(middleware.SetHeader("Access-Control-Allow-Origin", "*"))
	router.Use(middleware.SetHeader("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS"))
	router.Use(middleware.SetHeader("Access-Control-Allow-Headers", "Content-Type, Authorization"))

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