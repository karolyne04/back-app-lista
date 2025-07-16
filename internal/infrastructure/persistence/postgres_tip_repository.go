package persistence

import (
	"context"
	"database/sql"
	"karol/internal/domain/entity"
	"karol/internal/domain/repository"
)

type PostgresTipRepository struct {
	db *sql.DB
}

func NewPostgresTipRepository(db *sql.DB) repository.TipRepository {
	return &PostgresTipRepository{db: db}
}

func (r *PostgresTipRepository) List(ctx context.Context) ([]*entity.Tip, error) {
	query := `SELECT id, text FROM tips`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tips []*entity.Tip
	for rows.Next() {
		var t entity.Tip
		if err := rows.Scan(&t.ID, &t.Text); err != nil {
			return nil, err
		}
		tips = append(tips, &t)
	}
	return tips, nil
} 