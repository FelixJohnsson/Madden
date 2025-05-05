package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/madn-optimo/backend/repositories"
	"github.com/rs/zerolog/log"
)

type SaleHandler struct {
    repo *repositories.SaleRepository
}

func NewSaleHandler(repo *repositories.SaleRepository) *SaleHandler {
    return &SaleHandler{repo: repo}
}

func (h *SaleHandler) GetAll(w http.ResponseWriter, r *http.Request) {
    sales, err := h.repo.GetAll()
    if err != nil {
        log.Error().Err(err).Msg("Failed to get sales")
        http.Error(w, "Failed to fetch sales", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(sales)
}

func (h *SaleHandler) GetByID(w http.ResponseWriter, r *http.Request) {
    idStr := chi.URLParam(r, "id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        log.Error().Err(err).Msg("Invalid sale ID")
        http.Error(w, "Invalid sale ID", http.StatusBadRequest)
        return
    }

    sale, err := h.repo.GetByID(id)
    if err != nil {
        log.Error().Err(err).Msg("Failed to fetch sale")
        http.Error(w, "Failed to fetch sale", http.StatusInternalServerError)
        return
    }
    if sale == nil {
        log.Error().Msg("Sale not found")
        http.Error(w, "Sale not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(sale)
}
