package usecase

import (
	"context"
	"karol/internal/domain/entity"
	"karol/internal/domain/repository"
)

type TipUseCase struct {
	tipRepo repository.TipRepository
}

func NewTipUseCase(repo repository.TipRepository) *TipUseCase {
	return &TipUseCase{tipRepo: repo}
}

func (uc *TipUseCase) ListTips(ctx context.Context) ([]*entity.Tip, error) {
	return uc.tipRepo.List(ctx)
} 