package persistence

import (
	"context"
	"database/sql"
	"karol/internal/domain/entity"
	"karol/internal/domain/repository"
)

type PostgresCategoryRepository struct {
	db *sql.DB
}

func NewPostgresCategoryRepository(db *sql.DB) repository.CategoryRepository {
	return &PostgresCategoryRepository{db: db}
}

func (r *PostgresCategoryRepository) List(ctx context.Context) ([]*entity.Category, error) {
	query := `SELECT id, name FROM categories`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var categories []*entity.Category
	for rows.Next() {
		var c entity.Category
		if err := rows.Scan(&c.ID, &c.Name); err != nil {
			return nil, err
		}
		categories = append(categories, &c)
	}
	return categories, nil
} 