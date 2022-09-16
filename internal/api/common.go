package api

import (
	"github.com/gignhit/teslalabz/internal/model"
	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2"
)

type Handlers interface {
	Routers(app fiber.Router, authHandler fiber.Handler, middlewares ...fiber.Handler)
	DB() *pg.DB
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	resp := err.(*model.ErrorResponse)
	c.SendStatus(resp.StatusCode)
	return c.JSON(resp)
}
