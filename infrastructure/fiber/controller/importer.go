package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/middleware"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/request"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/response"
	"github.com/team-inu/inu-backyard/internal/validator"
	"github.com/team-inu/inu-backyard/usecase"
)

type importerController struct {
	importerUseCase usecase.ImporterUseCase
	validator       validator.PayloadValidator
}

func NewImporterController(validator validator.PayloadValidator, importerUseCase usecase.ImporterUseCase) importerController {
	return importerController{
		importerUseCase: importerUseCase,
		validator:       validator,
	}
}

func (c importerController) Import(ctx *fiber.Ctx) error {
	var payload request.ImportCoursePayload
	if ok, err := c.validator.Validate(&payload, ctx); !ok {
		return err
	}

	user := middleware.GetUserFromCtx(ctx)

	err := c.importerUseCase.UpdateOrCreate(
		payload.CourseId,
		user.Id,
		payload.StudentIds,
		payload.CourseLearningOutcomes,
		payload.AssignmentGroups,
		false,
	)

	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, nil)
}
