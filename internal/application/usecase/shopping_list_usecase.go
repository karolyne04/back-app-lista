package usecase

import (
	"context"
	"karol/internal/domain/entity"
	"karol/internal/domain/repository"
)

type ShoppingListUseCase struct {
	shoppingListRepo repository.ShoppingListRepository
}

func NewShoppingListUseCase(repo repository.ShoppingListRepository) *ShoppingListUseCase {
	return &ShoppingListUseCase{shoppingListRepo: repo}
}

func (uc *ShoppingListUseCase) CreateList(ctx context.Context, list *entity.ShoppingList) error {
	return uc.shoppingListRepo.Save(ctx, list)
}

func (uc *ShoppingListUseCase) GetListsByUser(ctx context.Context, userID string) ([]*entity.ShoppingList, error) {
	return uc.shoppingListRepo.FindByUserID(ctx, userID)
}

func (uc *ShoppingListUseCase) UpdateList(ctx context.Context, list *entity.ShoppingList) error {
	return uc.shoppingListRepo.Update(ctx, list)
}

func (uc *ShoppingListUseCase) DeleteList(ctx context.Context, listID string) error {
	return uc.shoppingListRepo.Delete(ctx, listID)
}

func (uc *ShoppingListUseCase) FinishList(ctx context.Context, listID string) error {
	return uc.shoppingListRepo.Finish(ctx, listID)
} 