package repository

import (
	"context"
	"karol/internal/domain/entity"
)

type ListItemRepository interface {
	Add(ctx context.Context, item *entity.ListItem) error
	Update(ctx context.Context, item *entity.ListItem) error
	Remove(ctx context.Context, id string) error
	FindByID(ctx context.Context, id string) (*entity.ListItem, error)
	FindByListID(ctx context.Context, listID string) ([]*entity.ListItem, error)
	Check(ctx context.Context, id string) error
} 