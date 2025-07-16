package persistence

import (
	"context"
	"database/sql"
	"karol/internal/domain/entity"
	"karol/internal/domain/repository"
)

type PostgresListItemRepository struct {
	db *sql.DB
}

func NewPostgresListItemRepository(db *sql.DB) repository.ListItemRepository {
	return &PostgresListItemRepository{db: db}
}

func (r *PostgresListItemRepository) Add(ctx context.Context, item *entity.ListItem) error {
	query := `INSERT INTO list_items (id, shopping_list_id, name, quantity, category_id, status, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.ExecContext(ctx, query, item.ID, item.ShoppingListID, item.Name, item.Quantity, item.CategoryID, item.Status, item.CreatedAt)
	return err
}
func (r *PostgresListItemRepository) Update(ctx context.Context, item *entity.ListItem) error {
	query := `UPDATE list_items SET name = $1, quantity = $2, category_id = $3, status = $4 WHERE id = $5`
	_, err := r.db.ExecContext(ctx, query, item.Name, item.Quantity, item.CategoryID, item.Status, item.ID)
	return err
}
func (r *PostgresListItemRepository) Remove(ctx context.Context, id string) error {
	query := `DELETE FROM list_items WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
func (r *PostgresListItemRepository) FindByID(ctx context.Context, id string) (*entity.ListItem, error) {
	query := `SELECT id, shopping_list_id, name, quantity, category_id, status, created_at FROM list_items WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)
	var i entity.ListItem
	if err := row.Scan(&i.ID, &i.ShoppingListID, &i.Name, &i.Quantity, &i.CategoryID, &i.Status, &i.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &i, nil
}
func (r *PostgresListItemRepository) FindByListID(ctx context.Context, listID string) ([]*entity.ListItem, error) {
	query := `SELECT id, shopping_list_id, name, quantity, category_id, status, created_at FROM list_items WHERE shopping_list_id = $1`
	rows, err := r.db.QueryContext(ctx, query, listID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*entity.ListItem
	for rows.Next() {
		var i entity.ListItem
		if err := rows.Scan(&i.ID, &i.ShoppingListID, &i.Name, &i.Quantity, &i.CategoryID, &i.Status, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	return items, nil
}
func (r *PostgresListItemRepository) Check(ctx context.Context, id string) error {
	query := `UPDATE list_items SET status = 'comprado' WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
} 