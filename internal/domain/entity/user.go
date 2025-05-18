package entity

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Erros específicos do domínio
var (
	ErrUsernameTaken = errors.New("nome de usuário já está em uso")
	ErrEmailTaken    = errors.New("email já está em uso")
)

// User representa a entidade de domínio para usuários
type User struct {
	ID        string
	Username  string
	Password  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewUser cria uma nova instância de User
func NewUser(username, password, email string) (*User, error) {
	if username == "" {
		return nil, errors.New("nome de usuário não pode ser vazio")
	}
	if password == "" {
		return nil, errors.New("senha não pode ser vazia")
	}
	if email == "" {
		return nil, errors.New("email não pode ser vazio")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		Username:  username,
		Password:  string(hashedPassword),
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

// ValidatePassword verifica se a senha fornecida corresponde à senha do usuário
func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
