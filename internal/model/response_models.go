package model

import (
	"github.com/gignhit/teslalabz/internal/entity"
)

type GetCategoryResponse struct {
	Id      int              `json:"id"`
	Name    string           `json:"name"`
	Liquids []CategoryLiquid `json:"liquids"`
}

type CategoryLiquid struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

type GetCommentsResponse struct {
	LiquidId int              `json:"liquid_id"`
	Comments []entity.Comment `json:"comments"`
}

type TokenResponse struct {
	Token string `json:"token"`
}
