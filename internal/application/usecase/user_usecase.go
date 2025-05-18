package usecase

import (
	"context"

	"karol/internal/domain/entity"
	"karol/internal/domain/repository"
)

// UserUseCase implementa os casos de uso relacionados a usuários
type UserUseCase struct {
	userRepository repository.UserRepository
}

// NewUserUseCase cria uma nova instância de UserUseCase
func NewUserUseCase(userRepository repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepository: userRepository,
	}
}

// CreateUser cria um novo usuário
func (uc *UserUseCase) CreateUser(ctx context.Context, user *entity.User) error {
	// Verificar se já existe um usuário com o mesmo nome de usuário
	existingUser, err := uc.userRepository.FindByUsername(ctx, user.Username)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return entity.ErrUsernameTaken
	}

	// Verificar se já existe um usuário com o mesmo email
	existingUser, err = uc.userRepository.FindByEmail(ctx, user.Email)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return entity.ErrEmailTaken
	}

	// Salvar usuário
	return uc.userRepository.Save(ctx, user)
}

// GetUserByID busca um usuário pelo ID
func (uc *UserUseCase) GetUserByID(ctx context.Context, id string) (*entity.User, error) {
	return uc.userRepository.FindByID(ctx, id)
}

// UpdateUser atualiza um usuário existente
func (uc *UserUseCase) UpdateUser(ctx context.Context, user *entity.User) error {
	return uc.userRepository.Update(ctx, user)
}

// DeleteUser remove um usuário
func (uc *UserUseCase) DeleteUser(ctx context.Context, id string) error {
	return uc.userRepository.Delete(ctx, id)
}