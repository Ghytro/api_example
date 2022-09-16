package liquid

import (
	"context"

	"github.com/gignhit/teslalabz/internal/entity"
	"github.com/gignhit/teslalabz/internal/repository"
)

// todo: перенести всю бизнес логику из api в service

type Service struct {
	repo *repository.LiquidRepository
}

func NewService(liquidRepo *repository.LiquidRepository) *Service {
	return &Service{repo: liquidRepo}
}

func (s *Service) GetLiquid(ctx context.Context, liquidId int) (*entity.Liquid, error) {
	return s.repo.GetLiquid(ctx, liquidId)
}

func (s *Service) GetLiquidPrice(ctx context.Context, liquidId int, strength string, volume int, dopping string) (*entity.Price, error) {
	return s.repo.GetLiquidPrice(ctx, liquidId, strength, volume, dopping)
}
