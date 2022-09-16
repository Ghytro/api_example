package orders

import (
	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2"
)

type OrdersApi struct {
	db *pg.DB
}

func NewOrdersApi(db *pg.DB) *OrdersApi {
	return &OrdersApi{db}
}

func (a *OrdersApi) DB() *pg.DB {
	return a.db
}

func (a *OrdersApi) Routers(app fiber.Router, authHandler fiber.Handler, middlewares ...fiber.Handler) {

}
