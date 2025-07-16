package usecase

import (
	"context"
	"karol/internal/domain/entity"
	"karol/internal/domain/repository"
)

type ListItemUseCase struct {
	listItemRepo repository.ListItemRepository
}

func NewListItemUseCase(repo repository.ListItemRepository) *ListItemUseCase {
	return &ListItemUseCase{listItemRepo: repo}
}

func (uc *ListItemUseCase) AddItem(ctx context.Context, item *entity.ListItem) error {
	return uc.listItemRepo.Add(ctx, item)
}

func (uc *ListItemUseCase) UpdateItem(ctx context.Context, item *entity.ListItem) error {
	return uc.listItemRepo.Update(ctx, item)
}

func (uc *ListItemUseCase) RemoveItem(ctx context.Context, itemID string) error {
	return uc.listItemRepo.Remove(ctx, itemID)
}

func (uc *ListItemUseCase) CheckItem(ctx context.Context, itemID string) error {
	return uc.listItemRepo.Check(ctx, itemID)
} 