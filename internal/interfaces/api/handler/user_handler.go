package handler

import (
	"encoding/json"
	"net/http"

	"karol/internal/application/usecase"
	"karol/internal/domain/entity"
)

// UserHandler gerencia as requisições relacionadas a usuários
type UserHandler struct {
	userUseCase *usecase.UserUseCase
}

// CreateUserRequest representa os dados para criação de um usuário
type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// UserResponse representa a resposta com dados do usuário
type UserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// NewUserHandler cria uma nova instância de UserHandler
func NewUserHandler(userUseCase *usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

// Create processa a requisição de criação de usuário
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	// Verificar método HTTP
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	// Decodificar requisição
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Requisição inválida", http.StatusBadRequest)
		return
	}

	// Criar usuário
	user, err := entity.NewUser(req.Username, req.Password, req.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Salvar usuário
	if err := h.userUseCase.CreateUser(r.Context(), user); err != nil {
		http.Error(w, "Erro ao criar usuário: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Preparar resposta
	resp := UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	// Enviar resposta
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

// GetProfile retorna o perfil do usuário autenticado
func (h *UserHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	// Verificar método HTTP
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	// Obter ID do usuário do contexto
	userID, ok := r.Context().Value("userID").(string)
	if !ok {
		http.Error(w, "Usuário não autenticado", http.StatusUnauthorized)
		return
	}

	// Buscar usuário
	user, err := h.userUseCase.GetUserByID(r.Context(), userID)
	if err != nil {
		http.Error(w, "Erro ao buscar usuário: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if user == nil {
		http.Error(w, "Usuário não encontrado", http.StatusNotFound)
		return
	}

	// Preparar resposta
	resp := UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	// Enviar resposta
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}