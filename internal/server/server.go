package server

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gignhit/teslalabz/internal/api"
	"github.com/gignhit/teslalabz/internal/api/users"
	"github.com/gignhit/teslalabz/internal/config"
	"github.com/gignhit/teslalabz/internal/entity"
	"github.com/gignhit/teslalabz/internal/model"
	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2"
)

type GracefulShutdownedServer struct {
	app *fiber.App
}

func validateToken(token string) bool {
	if len(token) != 50 {
		return false
	}
	for _, c := range token {
		if !(c >= '0' && c <= '9' || c >= 'a' && c <= 'z') {
			return false
		}
	}
	return true
}

func putUserEntityByToken(usersApi *users.UsersApi) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authedUser := new(entity.AuthedUser)
		authedUser.Token = c.Get("X-Access-Token")
		if !validateToken(authedUser.Token) {
			return model.GenErrorResponse(
				errors.New("некорректный токен"),
				fiber.StatusUnauthorized,
			)
		}
		if err := usersApi.DB().ModelContext(c.Context(), &authedUser).
			Where("token = ?", authedUser.Token).
			Relation("User").
			Select(); err != nil {
			if err == pg.ErrNoRows {
				return model.GenErrorResponse(
					errors.New("некорректный токен"),
					fiber.StatusUnauthorized,
				)
			}
			return model.GenErrorResponse(err, fiber.StatusInternalServerError)
		}
		c.Locals("user_entity", authedUser.User)
		return c.Next()
	}
}

func NewGracefulShutdownServer(routers ...api.Handlers) *GracefulShutdownedServer {
	app := fiber.New(fiber.Config{
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
		IdleTimeout:  config.IdleTimeout,
		BodyLimit:    config.BodyLimit,
		ErrorHandler: api.ErrorHandler,
	})

	g := app.Group("/api/v1")
	g.Use(func(c *fiber.Ctx) error {
		c.Set("Content-Type", "application/json; charset=utf-8")
		return c.Next()
	})
	uApi := users.NewUsersApi(routers[0].DB())

	for _, route := range routers {
		route.Routers(g, putUserEntityByToken(uApi))
	}
	return &GracefulShutdownedServer{
		app: app,
	}
}

func (s *GracefulShutdownedServer) Listen() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		if err := s.app.Listen(config.Addr); err != nil {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Printf("Server is listenning at %s", config.Addr)
	<-done
	log.Println("Gracefully shutting down server...")
	if err := s.app.Shutdown(); err != nil {
		log.Fatalf("Graceful shutdown failed: %+v; forcing shutdown", err)
	}
	log.Println("Server shutted down")
}
