package handler

import (
	"encoding/json"
	"net/http"
	"karol/internal/application/usecase"
	"karol/internal/domain/entity"
)

type ListItemHandler struct {
	useCase *usecase.ListItemUseCase
}

func NewListItemHandler(useCase *usecase.ListItemUseCase) *ListItemHandler {
	return &ListItemHandler{useCase: useCase}
}

func (h *ListItemHandler) Add(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	var req entity.ListItem
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Requisição inválida", http.StatusBadRequest)
		return
	}
	if err := h.useCase.AddItem(r.Context(), &req); err != nil {
		http.Error(w, "Erro ao adicionar item: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *ListItemHandler) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	var req entity.ListItem
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Requisição inválida", http.StatusBadRequest)
		return
	}
	if err := h.useCase.UpdateItem(r.Context(), &req); err != nil {
		http.Error(w, "Erro ao atualizar item: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *ListItemHandler) Remove(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id obrigatório", http.StatusBadRequest)
		return
	}
	if err := h.useCase.RemoveItem(r.Context(), id); err != nil {
		http.Error(w, "Erro ao remover item: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *ListItemHandler) Check(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id obrigatório", http.StatusBadRequest)
		return
	}
	if err := h.useCase.CheckItem(r.Context(), id); err != nil {
		http.Error(w, "Erro ao marcar item como comprado: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
} 