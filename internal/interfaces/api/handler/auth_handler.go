package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"karol/internal/application/usecase"
)

// AuthHandler gerencia as requisições de autenticação
type AuthHandler struct {
	authUseCase *usecase.AuthUseCase
}

// LoginRequest representa os dados de requisição de login
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse representa a resposta de login bem-sucedido
type LoginResponse struct {
	Token     string `json:"token"`
	ExpiresIn int64  `json:"expires_in"`
}

// NewAuthHandler cria uma nova instância de AuthHandler
func NewAuthHandler(authUseCase *usecase.AuthUseCase) *AuthHandler {
	return &AuthHandler{
		authUseCase: authUseCase,
	}
}

// Login processa a requisição de login
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	// Verificar método HTTP
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	// Decodificar requisição
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Requisição inválida", http.StatusBadRequest)
		return
	}

	// Autenticar usuário
	token, err := h.authUseCase.Login(r.Context(), req.Username, req.Password)
	if err != nil {
		http.Error(w, "Credenciais inválidas", http.StatusUnauthorized)
		return
	}

	// Preparar resposta
	expiresIn := time.Hour * 24 // 24 horas
	resp := LoginResponse{
		Token:     token,
		ExpiresIn: int64(expiresIn.Seconds()),
	}

	// Enviar resposta
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}