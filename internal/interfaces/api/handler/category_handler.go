package handler

import (
	"encoding/json"
	"net/http"
	"karol/internal/application/usecase"
)

type CategoryHandler struct {
	useCase *usecase.CategoryUseCase
}

func NewCategoryHandler(useCase *usecase.CategoryUseCase) *CategoryHandler {
	return &CategoryHandler{useCase: useCase}
}

func (h *CategoryHandler) List(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	categories, err := h.useCase.ListCategories(r.Context())
	if err != nil {
		http.Error(w, "Erro ao buscar categorias: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
} 