package handler

import (
	"encoding/json"
	"net/http"
	"karol/internal/application/usecase"
)

type TipHandler struct {
	useCase *usecase.TipUseCase
}

func NewTipHandler(useCase *usecase.TipUseCase) *TipHandler {
	return &TipHandler{useCase: useCase}
}

func (h *TipHandler) List(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	tips, err := h.useCase.ListTips(r.Context())
	if err != nil {
		http.Error(w, "Erro ao buscar dicas: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tips)
} 