package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"

	"github.com/madn-optimo/backend/models"
	"github.com/madn-optimo/backend/repositories"
)

type PurchaseOrderHandler struct {
	repo *repositories.PurchaseOrderRepository
}

func NewPurchaseOrderHandler(repo *repositories.PurchaseOrderRepository) *PurchaseOrderHandler {
	return &PurchaseOrderHandler{repo: repo}
}

func (h *PurchaseOrderHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	orders, err := h.repo.GetAll()
	if err != nil {
		log.Error().Err(err).Msg("Failed to get purchase orders")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

func (h *PurchaseOrderHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	order, err := h.repo.GetByID(id)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get purchase order by ID")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if order == nil {
		http.Error(w, "Purchase order not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

func (h *PurchaseOrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	var order models.PurchaseOrder
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if order.CreatedAt.IsZero() {
			order.CreatedAt = time.Now()
	}
	order.Amount = int(order.Amount)

	id, err := h.repo.Create(order)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create purchase order")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

func (h *PurchaseOrderHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := h.repo.Delete(id); err != nil {
		log.Error().Err(err).Msg("Failed to delete purchase order")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
} 