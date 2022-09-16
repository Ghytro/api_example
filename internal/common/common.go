package common

import (
	"errors"
	"regexp"
	"strconv"

	"github.com/gignhit/teslalabz/internal/model"
	"github.com/gofiber/fiber/v2"
)

func GetIntIdFromContext(c *fiber.Ctx) (int, error) {
	result, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return 0, model.GenErrorResponse(
			errors.New(
				"Передан некорректный формат id",
			),
			fiber.StatusBadRequest,
		)
	}
	return int(result), err
}

func DecodeJsonModel(c *fiber.Ctx, req interface{}) error {
	if err := c.BodyParser(req); err != nil {
		return model.GenErrorResponse(
			errors.New(
				"Не получилось декодировать JSON. Проверьте, что Content-Type установлен в application/json",
			),
			fiber.StatusBadRequest,
		)
	}
	return nil
}

func CheckStringWithRegexp(s string, regex *regexp.Regexp) bool {
	f := regex.FindIndex([]byte(s))
	if f == nil {
		return false
	}
	return f[0] == 0 && f[1] == len(s)
}
