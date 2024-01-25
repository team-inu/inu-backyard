package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team-inu/inu-backyard/entity"
	"github.com/team-inu/inu-backyard/internal/validator"
)

func NewAuthMiddleware(
	validator validator.PayloadValidator,
	authUsecase entity.AuthUseCase,
) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		sid, err := validator.ValidateAuth(ctx)
		if sid == "" {
			return err
		}

		user, err := authUsecase.Authenticate(sid)
		if err != nil {
			return err
		}

		ctx.Locals("user", user)

		return ctx.Next()
	}
}

func GetUserFromCtx(ctx *fiber.Ctx) *entity.Lecturer {
	user, _ := ctx.Locals("user").(*entity.Lecturer)
	return user
}
