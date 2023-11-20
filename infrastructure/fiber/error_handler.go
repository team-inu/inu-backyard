package fiber

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	errs "github.com/team-inu/inu-backyard/entity/error"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/response"
	"go.uber.org/zap"
)

func errorHandler(logger *zap.Logger) fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		resStatus := fiber.StatusInternalServerError

		var domainError *errs.DomainError
		if errors.As(err, &domainError) {
			status, ok := response.DomainErrCodeToHttpStatus[domainError.Code]
			if ok {
				resStatus = status
			}
		}

		if jsonErr := response.NewErrorResponse(ctx, resStatus, err); jsonErr != nil {
			logger.Error(
				"Cannot marshal json response",
				zap.NamedError("json_error", jsonErr),
				zap.NamedError("error", err),
			)
			return ctx.Status(500).SendString(err.Error())
		}

		return nil
	}
}
