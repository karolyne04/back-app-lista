package repository

import (
	"context"
	"karol/internal/domain/entity"
)

type ShoppingListRepository interface {
	Save(ctx context.Context, list *entity.ShoppingList) error
	FindByID(ctx context.Context, id string) (*entity.ShoppingList, error)
	FindByUserID(ctx context.Context, userID string) ([]*entity.ShoppingList, error)
	Update(ctx context.Context, list *entity.ShoppingList) error
	Delete(ctx context.Context, id string) error
	Finish(ctx context.Context, id string) error
} 