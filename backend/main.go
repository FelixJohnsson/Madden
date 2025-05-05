package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/madn-optimo/backend/db"
	"github.com/madn-optimo/backend/handlers"
	"github.com/madn-optimo/backend/repositories"
	"github.com/madn-optimo/backend/routes"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	database, err := db.InitDB()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize database")
	}
	defer database.Close()

	saleRepo := repositories.NewSaleRepository(database)
	purchaseOrderRepo := repositories.NewPurchaseOrderRepository(database)

	saleHandler := handlers.NewSaleHandler(saleRepo)
	purchaseOrderHandler := handlers.NewPurchaseOrderHandler(purchaseOrderRepo)

	router := chi.NewRouter()
	routes.SetupRoutes(router, saleHandler, purchaseOrderHandler)

	addr := fmt.Sprintf(":%s", port)
	log.Info().Str("address", addr).Msg("Server starting")
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatal().Err(err).Msg("Server failed to start")
	}
}