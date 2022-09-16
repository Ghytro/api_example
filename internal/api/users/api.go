package users

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"regexp"
	"strings"

	"github.com/gignhit/teslalabz/internal/common"
	"github.com/gignhit/teslalabz/internal/entity"
	"github.com/gignhit/teslalabz/internal/model"
	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2"
)

type UsersApi struct {
	db *pg.DB
}

func NewUsersApi(db *pg.DB) *UsersApi {
	return &UsersApi{db}
}

func (a *UsersApi) DB() *pg.DB {
	return a.db
}

func (a *UsersApi) signIn(c *fiber.Ctx) error {
	var loginReq model.SignInRequest
	if err := common.DecodeJsonModel(c, &loginReq); err != nil {
		return err
	}
	var authed entity.AuthedUser
	if err := a.db.ModelContext(c.Context(), &authed).
		Where(
			"login = ? AND password = crypt(?, password)",
			loginReq.Login,
			loginReq.Password,
		).
		Select(); err != nil {
		if err == pg.ErrNoRows {
			return model.GenErrorResponse(
				errors.New("неверный логин или пароль"),
				fiber.StatusBadRequest,
			)
		}
		return model.GenErrorResponse(err, fiber.StatusInternalServerError)
	}
	return c.JSON(model.TokenResponse{Token: authed.Token})
}

func validateSignUpRequest(req *model.SignUpRequest) error {
	if req.Contacts.Phone == "" {
		return model.GenErrorResponse(
			errors.New("Не указан телефон"),
			fiber.StatusBadRequest,
		)
	}
	phoneRegex, err := regexp.Compile(`^((8|\+7)[\- ]?)?(\(?\d{3}\)?[\- ]?)?[\d\- ]{7,10}$`)
	if err != nil {
		return model.GenErrorResponse(
			errors.New("debug: incorrect regex for phone"),
			fiber.StatusInternalServerError,
		)
	}
	if !common.CheckStringWithRegexp(req.Contacts.Phone, phoneRegex) {
		return model.GenErrorResponse(
			fmt.Errorf("некорректный формат номера телефона: %s", req.Contacts.Phone),
			fiber.StatusBadRequest,
		)
	}
	if len(req.Password) < 6 {
		return model.GenErrorResponse(
			errors.New("длина пароля должна быть как минимум 6"),
			fiber.StatusBadRequest,
		)
	}

	var formattedPhone strings.Builder
	for _, c := range req.Contacts.Phone {
		if c >= '0' && c <= '9' {
			formattedPhone.WriteRune(c)
		}
	}
	req.Contacts.Phone = formattedPhone.String()
	return nil
}

func generateToken() string {
	allowed := "qwertyuiopasdfghjklzxcvbnm0123456789"
	var token strings.Builder
	for i := 0; i < 50; i++ {
		idx, _ := rand.Int(rand.Reader, big.NewInt(int64(len(allowed))))
		token.WriteByte(allowed[idx.Int64()])
	}
	return token.String()
}

func (a *UsersApi) signUp(c *fiber.Ctx) error {
	var req *model.SignUpRequest
	if err := common.DecodeJsonModel(c, req); err != nil {
		return err
	}
	if err := validateSignUpRequest(req); err != nil {
		return err
	}

	if err := a.db.ModelContext(c.Context(), (*entity.AuthedUser)(nil)).
		Where("login = ?", req.Login).
		Select(); err == nil {
		return model.GenErrorResponse(
			fmt.Errorf("уже есть пользователь с таким логином: %s"),
			fiber.StatusBadRequest,
		)
	}
	insertedUser := &entity.User{
		Name:   req.Name,
		Phone:  req.Contacts.Phone,
		VkLink: req.Contacts.VkLink,
		TgLink: req.Contacts.TgLink,
	}
	if _, err := a.db.ModelContext(c.Context(), insertedUser).Insert(); err != nil {
		return model.GenErrorResponse(err, fiber.StatusInternalServerError)
	}
	authedUser := &entity.AuthedUser{
		Login:    req.Login,
		Password: req.Password,
		Token:    generateToken(),
	}
	if _, err := a.db.ModelContext(c.Context(), authedUser).Insert(); err != nil {
		return model.GenErrorResponse(err, fiber.StatusInternalServerError)
	}
	return c.JSON(model.TokenResponse{Token: authedUser.Token})
}

func (a *UsersApi) Routers(app fiber.Router, authHandler fiber.Handler, middlewares ...fiber.Handler) {
	r := fiber.New()

	r.Get("/auth", a.signIn)
	r.Post("/auth", a.signUp)
}
