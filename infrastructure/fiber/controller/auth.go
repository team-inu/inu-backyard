package controller

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/team-inu/inu-backyard/entity"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/middleware"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/request"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/response"
	"github.com/team-inu/inu-backyard/internal/config"
	"github.com/team-inu/inu-backyard/internal/validator"
)

type AuthController struct {
	config          config.AuthConfig
	validator       validator.PayloadValidator
	authUseCase     entity.AuthUseCase
	lecturerUseCase entity.LecturerUseCase
}

func NewAuthController(
	config config.AuthConfig,
	authUseCase entity.AuthUseCase,
	lecturerUseCase entity.LecturerUseCase,
) *AuthController {
	return &AuthController{
		config:          config,
		validator:       validator.NewPayloadValidator(),
		authUseCase:     authUseCase,
		lecturerUseCase: lecturerUseCase,
	}
}

func (c AuthController) Me(ctx *fiber.Ctx) error {
	return response.NewSuccessResponse(ctx, fiber.StatusOK, middleware.GetUserFromCtx(ctx))
}

func (c AuthController) SignIn(ctx *fiber.Ctx) error {
	var payload request.SignInPayload
	if ok, err := c.validator.Validate(&payload, ctx); !ok {
		return err
	}

	ipAddress := ctx.IP()
	userAgent := ctx.Context().UserAgent()

	cookie, err := c.authUseCase.SignIn(payload.Email, payload.Password, ipAddress, string(userAgent))
	if err != nil {
		return err
	}
	ctx.Cookie(cookie)

	return response.NewSuccessResponse(ctx, fiber.StatusOK, fiber.Map{
		"expired_at": cookie.Expires,
	})
}

func (c AuthController) SignOut(ctx *fiber.Ctx) error {
	fmt.Println(c.config.Session.CookieName)

	sid := ctx.Cookies(c.config.Session.CookieName)
	cookie, err := c.authUseCase.SignOut(sid)
	if err != nil {
		return err
	}
	ctx.Cookie(cookie)

	return response.NewSuccessResponse(ctx, fiber.StatusOK, fiber.Map{
		"signout_at": time.Now(),
	})
}
