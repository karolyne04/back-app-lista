package persistence

import (
	"context"
	"database/sql"
	"karol/internal/domain/entity"
	"karol/internal/domain/repository"
	"time"
)

type PostgresShoppingListRepository struct {
	db *sql.DB
}

func NewPostgresShoppingListRepository(db *sql.DB) repository.ShoppingListRepository {
	return &PostgresShoppingListRepository{db: db}
}

func (r *PostgresShoppingListRepository) Save(ctx context.Context, list *entity.ShoppingList) error {
	query := `INSERT INTO shopping_lists (id, user_id, name, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.ExecContext(ctx, query, list.ID, list.UserID, list.Name, list.CreatedAt, list.UpdatedAt)
	return err
}

func (r *PostgresShoppingListRepository) FindByUserID(ctx context.Context, userID string) ([]*entity.ShoppingList, error) {
	query := `SELECT id, user_id, name, created_at, updated_at, finished_at FROM shopping_lists WHERE user_id = $1`
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var lists []*entity.ShoppingList
	for rows.Next() {
		var l entity.ShoppingList
		var finishedAt sql.NullTime
		if err := rows.Scan(&l.ID, &l.UserID, &l.Name, &l.CreatedAt, &l.UpdatedAt, &finishedAt); err != nil {
			return nil, err
		}
		if finishedAt.Valid {
			l.FinishedAt = &finishedAt.Time
		}
		lists = append(lists, &l)
	}
	return lists, nil
}
func (r *PostgresShoppingListRepository) FindByID(ctx context.Context, id string) (*entity.ShoppingList, error) {
	query := `SELECT id, user_id, name, created_at, updated_at, finished_at FROM shopping_lists WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)
	var l entity.ShoppingList
	var finishedAt sql.NullTime
	if err := row.Scan(&l.ID, &l.UserID, &l.Name, &l.CreatedAt, &l.UpdatedAt, &finishedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	if finishedAt.Valid {
		l.FinishedAt = &finishedAt.Time
	}
	return &l, nil
}
func (r *PostgresShoppingListRepository) Update(ctx context.Context, list *entity.ShoppingList) error {
	query := `UPDATE shopping_lists SET name = $1, updated_at = $2 WHERE id = $3`
	_, err := r.db.ExecContext(ctx, query, list.Name, list.UpdatedAt, list.ID)
	return err
}
func (r *PostgresShoppingListRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM shopping_lists WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
func (r *PostgresShoppingListRepository) Finish(ctx context.Context, id string) error {
	query := `UPDATE shopping_lists SET finished_at = $1 WHERE id = $2`
	_, err := r.db.ExecContext(ctx, query, time.Now(), id)
	return err
} 