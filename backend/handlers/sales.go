package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"

	"github.com/madn-optimo/backend/repositories"
)

type SaleHandler struct {
	repo *repositories.SaleRepository
}

func NewSaleHandler(repo *repositories.SaleRepository) *SaleHandler {
	return &SaleHandler{repo: repo}
}

func (h *SaleHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	page := 1
	pageSize := 10

	pageParam := r.URL.Query().Get("page")
	if pageParam != "" {
		pageInt, err := strconv.Atoi(pageParam)
		if err == nil && pageInt > 0 {
			page = pageInt
		}
	}

	pageSizeParam := r.URL.Query().Get("pageSize")
	if pageSizeParam != "" {
		pageSizeInt, err := strconv.Atoi(pageSizeParam)
		if err == nil && pageSizeInt > 0 && pageSizeInt <= 100 {
			pageSize = pageSizeInt
		}
	}

	sales, err := h.repo.GetAll(page, pageSize)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get sales")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sales)
}

func (h *SaleHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	sale, err := h.repo.GetByID(id)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get sale by ID")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if sale == nil {
		http.Error(w, "Sale not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sale)
} 