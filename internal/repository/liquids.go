package repository

import (
	"context"

	"github.com/gignhit/teslalabz/internal/entity"
)

type LiquidRepository struct {
	baseRepo
}

func (repo *LiquidRepository) GetLiquid(ctx context.Context, liquidId int, requiredRelations ...string) (*entity.Liquid, error) {
	liquid := &entity.Liquid{
		Id: liquidId,
	}
	q := repo.db.ModelContext(ctx, liquid).WherePK()
	for _, r := range requiredRelations {
		q = q.Relation(r)
	}
	if err := q.Select(); err != nil {
		return nil, err
	}
	return liquid, nil
}

func (repo *LiquidRepository) GetLiquidPrice(ctx context.Context, liquidId int, strength string, volume int, dopping string) (*entity.Price, error) {
	price := new(entity.Price)
	if err := repo.db.ModelContext(ctx, price).
		Where(
			"liquid_id = ? AND strength = ? AND volume = ? AND dopping = ?",
			liquidId,
			strength,
			volume,
			dopping,
		).
		Select(); err != nil {
		return nil, err
	}
	return price, nil
}
