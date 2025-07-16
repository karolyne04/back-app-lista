package usecase

import (
	"context"
	"karol/internal/domain/entity"
	"karol/internal/domain/repository"
)

type CategoryUseCase struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryUseCase(repo repository.CategoryRepository) *CategoryUseCase {
	return &CategoryUseCase{categoryRepo: repo}
}

func (uc *CategoryUseCase) ListCategories(ctx context.Context) ([]*entity.Category, error) {
	return uc.categoryRepo.List(ctx)
} 