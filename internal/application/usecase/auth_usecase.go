package usecase

import (
	"context"
	"errors"
	"time"

	"karol/internal/domain/repository"
	"karol/internal/domain/service"
)

// AuthUseCase implementa os casos de uso relacionados à autenticação
type AuthUseCase struct {
	userRepository repository.UserRepository
	jwtService     *service.JWTService
}

// NewAuthUseCase cria uma nova instância de AuthUseCase
func NewAuthUseCase(userRepository repository.UserRepository, jwtService *service.JWTService) *AuthUseCase {
	return &AuthUseCase{
		userRepository: userRepository,
		jwtService:     jwtService,
	}
}

// Login autentica um usuário e retorna um token JWT
func (uc *AuthUseCase) Login(ctx context.Context, username, password string) (string, error) {
	// Buscar usuário pelo nome de usuário
	user, err := uc.userRepository.FindByUsername(ctx, username)
	if err != nil {
		return "", err
	}

	// Verificar se o usuário existe
	if user == nil {
		return "", errors.New("usuário não encontrado")
	}

	// Validar senha
	if !user.ValidatePassword(password) {
		return "", errors.New("senha inválida")
	}

	// Gerar token JWT
	token, err := uc.jwtService.GenerateToken(user.ID, user.Username, 24*time.Hour)
	if err != nil {
		return "", err
	}

	return token, nil
}
