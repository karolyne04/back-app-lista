package middleware

import (
	"context"
	"net/http"
	"strings"

	"karol/internal/domain/service"
)

// AuthMiddleware é um middleware para autenticação JWT
type AuthMiddleware struct {
	jwtService *service.JWTService
}

// NewAuthMiddleware cria uma nova instância de AuthMiddleware
func NewAuthMiddleware(jwtService *service.JWTService) *AuthMiddleware {
	return &AuthMiddleware{
		jwtService: jwtService,
	}
}

// Authenticate verifica se o token JWT é válido
func (m *AuthMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extrair token do cabeçalho Authorization
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Autorização necessária", http.StatusUnauthorized)
			return
		}

		// Verificar formato do token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Formato de autorização inválido", http.StatusUnauthorized)
			return
		}

		// Validar token
		token := parts[1]
		claims, err := m.jwtService.ValidateToken(token)
		if err != nil {
			http.Error(w, "Token inválido: "+err.Error(), http.StatusUnauthorized)
			return
		}

		// Adicionar claims ao contexto da requisição
		ctx := context.WithValue(r.Context(), "userID", claims.UserID)
		ctx = context.WithValue(ctx, "username", claims.Username)

		// Continuar com o próximo handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}