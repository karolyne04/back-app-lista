package api

import (
	"net/http"

	"karol/internal/interfaces/api/handler"
	"karol/internal/interfaces/api/middleware"
)

// Router gerencia as rotas da API
type Router struct {
	authHandler    *handler.AuthHandler
	userHandler    *handler.UserHandler
	authMiddleware *middleware.AuthMiddleware
}

// NewRouter cria uma nova instância de Router
func NewRouter(
	authHandler *handler.AuthHandler,
	userHandler *handler.UserHandler,
	authMiddleware *middleware.AuthMiddleware,
) *Router {
	return &Router{
		authHandler:    authHandler,
		userHandler:    userHandler,
		authMiddleware: authMiddleware,
	}
}

// Setup configura as rotas da API
func (r *Router) Setup() http.Handler {
	mux := http.NewServeMux()

	// Rotas públicas
	mux.HandleFunc("/api/auth/login", r.authHandler.Login)
	mux.HandleFunc("/api/users", r.userHandler.Create)

	// Rotas protegidas
	protectedMux := http.NewServeMux()
	protectedMux.HandleFunc("/api/users/me", r.userHandler.GetProfile)
	
	// Aplicar middleware de autenticação às rotas protegidas
	mux.Handle("/api/users/me", r.authMiddleware.Authenticate(protectedMux))

	return mux
}