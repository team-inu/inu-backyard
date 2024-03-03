package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team-inu/inu-backyard/entity"
	"github.com/team-inu/inu-backyard/internal/validator"
)

func NewAuthMiddleware(
	validator validator.PayloadValidator,
	authUseCase entity.AuthUseCase,
) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		sid, err := validator.ValidateAuth(ctx)
		if sid == "" {
			return err
		}

		user, err := authUseCase.Authenticate(sid)
		if err != nil {
			return err
		}

		ctx.Locals("user", user)

		return ctx.Next()
	}
}

func GetUserFromCtx(ctx *fiber.Ctx) *entity.User {
	user, _ := ctx.Locals("user").(*entity.User)
	return user
}
