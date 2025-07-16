package handler

import (
	"encoding/json"
	"net/http"
	"karol/internal/application/usecase"
	"karol/internal/domain/entity"
	"strings"
)

type ShoppingListHandler struct {
	useCase *usecase.ShoppingListUseCase
}

func NewShoppingListHandler(useCase *usecase.ShoppingListUseCase) *ShoppingListHandler {
	return &ShoppingListHandler{useCase: useCase}
}

func (h *ShoppingListHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	var req entity.ShoppingList
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Requisição inválida", http.StatusBadRequest)
		return
	}
	if err := h.useCase.CreateList(r.Context(), &req); err != nil {
		http.Error(w, "Erro ao criar lista: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *ShoppingListHandler) List(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "user_id obrigatório", http.StatusBadRequest)
		return
	}
	lists, err := h.useCase.GetListsByUser(r.Context(), userID)
	if err != nil {
		http.Error(w, "Erro ao buscar listas: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lists)
}

func (h *ShoppingListHandler) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	var req entity.ShoppingList
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Requisição inválida", http.StatusBadRequest)
		return
	}
	if err := h.useCase.UpdateList(r.Context(), &req); err != nil {
		http.Error(w, "Erro ao atualizar lista: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *ShoppingListHandler) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id obrigatório", http.StatusBadRequest)
		return
	}
	if err := h.useCase.DeleteList(r.Context(), id); err != nil {
		http.Error(w, "Erro ao deletar lista: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *ShoppingListHandler) Finish(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id obrigatório", http.StatusBadRequest)
		return
	}
	if err := h.useCase.FinishList(r.Context(), id); err != nil {
		http.Error(w, "Erro ao finalizar lista: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
} 