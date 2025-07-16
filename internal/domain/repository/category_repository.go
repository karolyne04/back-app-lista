package repository

import (
	"context"
	"karol/internal/domain/entity"
)

type CategoryRepository interface {
	List(ctx context.Context) ([]*entity.Category, error)
} 