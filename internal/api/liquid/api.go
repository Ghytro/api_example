package liquid

import (
	"errors"
	"fmt"

	"github.com/gignhit/teslalabz/internal/common"
	"github.com/gignhit/teslalabz/internal/entity"
	"github.com/gignhit/teslalabz/internal/model"
	"github.com/gignhit/teslalabz/pkg/algorithm"
	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2"
)

type LiquidsApi struct {
	db *pg.DB
}

func NewLiquidsApi(db *pg.DB) *LiquidsApi {
	return &LiquidsApi{db}
}

func (a *LiquidsApi) DB() *pg.DB {
	return a.db
}

func (a *LiquidsApi) getLiquid(c *fiber.Ctx) error {
	liquidId, err := common.GetIntIdFromContext(c)
	if err != nil {
		return err
	}
	result := &entity.Liquid{
		Id: int(liquidId),
	}
	if err = a.db.ModelContext(c.Context(), result).
		WherePK().
		Select(); err != nil {
		if err != pg.ErrNoRows {
			return model.GenErrorResponse(err, fiber.StatusInternalServerError)
		}
		return model.GenErrorResponse(
			fmt.Errorf("Не найдено жидкости с id %d", liquidId),
			fiber.StatusBadRequest,
		)
	}
	return c.JSON(result)
}

func (a *LiquidsApi) getLiquidPrice(c *fiber.Ctx) error {
	var req *model.PriceRequest
	if err := common.DecodeJsonModel(c, req); err != nil {
		return err
	}
	var result *entity.Price
	if err := a.db.ModelContext(c.Context(), result).
		Where("liquid_id = ? AND strength = ? AND volume = ? AND dopping = ?").
		Select(); err != nil {
		if err == pg.ErrNoRows {
			return model.GenErrorResponse(
				errors.New("не найдено жидкости с указанными параметрами"),
				fiber.StatusBadRequest,
			)
		}
		return model.GenErrorResponse(err, fiber.StatusInternalServerError)
	}
	return c.JSON(result)
}

func (a *LiquidsApi) getCategories(c *fiber.Ctx) error {
	var categories []entity.Category
	err := a.db.ModelContext(c.Context(), &categories).Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return model.GenErrorResponse(
				errors.New("категорий нет"),
				fiber.StatusBadRequest,
			)
		}
		return model.GenErrorResponse(err, fiber.StatusInternalServerError)
	}
	return c.JSON(categories)
}

func (a *LiquidsApi) getCategory(c *fiber.Ctx) error {
	categoryId, err := common.GetIntIdFromContext(c)
	if err != nil {
		return err
	}

	var liquids []entity.Liquid
	err = a.db.ModelContext(c.Context(), &liquids).
		Where("category_id = ?", categoryId).
		Relation("Category").
		Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return model.GenErrorResponse(
				fmt.Errorf("категория с id не найдена: %d", categoryId),
				fiber.StatusBadRequest,
			)
		}
		return model.GenErrorResponse(err, fiber.StatusInternalServerError)
	}
	return c.JSON(
		model.GetCategoryResponse{
			Id:   categoryId,
			Name: liquids[0].Category.Name,
			Liquids: algorithm.Transformed(
				liquids,
				func(l entity.Liquid) model.CategoryLiquid {
					return model.CategoryLiquid{
						Id:    l.Id,
						Name:  l.Name,
						Image: l.Image,
					}
				}),
		},
	)
}

func (a *LiquidsApi) Routers(app fiber.Router, authHandler fiber.Handler, middlewares ...fiber.Handler) {
	r := fiber.New()
	r.Get("/:id", a.getLiquid)
	r.Get("/:id/price", a.getLiquidPrice)

	gCategories := r.Group("/categories")
	gCategories.Get("/:id", a.getCategory)
	gCategories.Get("/", a.getCategories)

	app.Mount("/liquids", r)
}
