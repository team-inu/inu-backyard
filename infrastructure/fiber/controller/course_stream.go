package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team-inu/inu-backyard/entity"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/request"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/response"
	"github.com/team-inu/inu-backyard/internal/validator"
)

type courseStreamController struct {
	CourseStreamUseCase entity.CourseStreamsUseCase
	Validator           validator.PayloadValidator
}

func NewCourseStreamController(validator validator.PayloadValidator, courseStreamUseCase entity.CourseStreamsUseCase) *courseStreamController {
	return &courseStreamController{
		CourseStreamUseCase: courseStreamUseCase,
		Validator:           validator,
	}
}

// 555
func (c courseStreamController) Get(ctx *fiber.Ctx) error {
	var payload request.GetCourseStreamRequestPayload
	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	err := payload.Validate()
	if err != nil {
		return err
	}

	if payload.TargetCourseId != "" {
		streamCourse, err := c.CourseStreamUseCase.GetByTargetCourseId(payload.TargetCourseId)
		if err != nil {
			return err
		}

		return response.NewSuccessResponse(ctx, fiber.StatusOK, streamCourse)
	}

	if payload.FromCourseId != "" {
		streamCourse, err := c.CourseStreamUseCase.GetByFromCourseId(payload.FromCourseId)
		if err != nil {
			return err
		}

		return response.NewSuccessResponse(ctx, fiber.StatusOK, streamCourse)
	}

	return nil
}

func (c courseStreamController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("courseStreamId")

	err := c.CourseStreamUseCase.Delete(id)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, nil)
}

func (c courseStreamController) Create(ctx *fiber.Ctx) error {
	var payload request.CreateCourseStreamRequestPayload
	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	err := c.CourseStreamUseCase.Create(
		payload.FromCourseId,
		payload.TargetCourseId,
		payload.StreamType,
		payload.Comment,
	)

	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusCreated, nil)
}
