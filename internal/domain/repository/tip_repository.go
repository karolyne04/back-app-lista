package repository

import (
	"context"
	"karol/internal/domain/entity"
)

type TipRepository interface {
	List(ctx context.Context) ([]*entity.Tip, error)
} 