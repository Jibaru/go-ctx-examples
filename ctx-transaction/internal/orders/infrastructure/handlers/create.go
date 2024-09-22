package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/jibaru/ctx-transaction/internal/orders/application"
)

type CreateOrderHandler struct {
	service application.CreateOrderService
}

func NewCreateOrderHandler(service application.CreateOrderService) *CreateOrderHandler {
	return &CreateOrderHandler{service: service}
}

func (h *CreateOrderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var input application.CreateOrderInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.Exec(r.Context(), input); err != nil {
		http.Error(w, "Failed to create order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
