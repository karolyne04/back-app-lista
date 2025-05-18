package persistence

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"karol/internal/domain/entity"
	"karol/internal/domain/repository"
)

// PostgresUserRepository implementa a interface UserRepository
type PostgresUserRepository struct {
	db *PostgresDB
}

// NewPostgresUserRepository cria uma nova instância de PostgresUserRepository
func NewPostgresUserRepository(db *PostgresDB) repository.UserRepository {
	return &PostgresUserRepository{db: db}
}

// Save persiste um usuário no banco de dados
func (r *PostgresUserRepository) Save(ctx context.Context, user *entity.User) error {
	query := `
		INSERT INTO users (id, username, password, email, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.db.DB.ExecContext(
		ctx,
		query,
		user.ID,
		user.Username,
		user.Password,
		user.Email,
		user.CreatedAt,
		user.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

// FindByID busca um usuário pelo ID
func (r *PostgresUserRepository) FindByID(ctx context.Context, id string) (*entity.User, error) {
	query := `
		SELECT id, username, password, email, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	var user entity.User
	err := r.db.DB.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

// FindByUsername busca um usuário pelo nome de usuário
func (r *PostgresUserRepository) FindByUsername(ctx context.Context, username string) (*entity.User, error) {
	query := `
		SELECT id, username, password, email, created_at, updated_at
		FROM users
		WHERE username = $1
	`

	var user entity.User
	err := r.db.DB.QueryRowContext(ctx, query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

// FindByEmail busca um usuário pelo email
func (r *PostgresUserRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	query := `
		SELECT id, username, password, email, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	var user entity.User
	err := r.db.DB.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

// Update atualiza um usuário existente
func (r *PostgresUserRepository) Update(ctx context.Context, user *entity.User) error {
	query := `
		UPDATE users
		SET username = $1, password = $2, email = $3, updated_at = $4
		WHERE id = $5
	`

	user.UpdatedAt = time.Now()

	_, err := r.db.DB.ExecContext(
		ctx,
		query,
		user.Username,
		user.Password,
		user.Email,
		user.UpdatedAt,
		user.ID,
	)

	return err
}

// Delete remove um usuário pelo ID
func (r *PostgresUserRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.DB.ExecContext(ctx, query, id)
	return err
}
