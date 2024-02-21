package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team-inu/inu-backyard/entity"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/request"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/response"
	"github.com/team-inu/inu-backyard/internal/validator"
)

type courseLearningOutcomeController struct {
	courseLearningOutcomeUsecase entity.CourseLearningOutcomeUsecase
	Validator                    validator.PayloadValidator
}

func NewCourseLearningOutcomeController(validator validator.PayloadValidator, courseLearningOutcomeUsecase entity.CourseLearningOutcomeUsecase) *courseLearningOutcomeController {
	return &courseLearningOutcomeController{
		courseLearningOutcomeUsecase: courseLearningOutcomeUsecase,
		Validator:                    validator,
	}
}

func (c courseLearningOutcomeController) GetAll(ctx *fiber.Ctx) error {
	clos, err := c.courseLearningOutcomeUsecase.GetAll()
	if err != nil {
		return err
	}

	return ctx.JSON(clos)
}

func (c courseLearningOutcomeController) GetById(ctx *fiber.Ctx) error {
	cloId := ctx.Params("cloId")

	clo, err := c.courseLearningOutcomeUsecase.GetById(cloId)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusCreated, clo)
}

func (c courseLearningOutcomeController) GetByCourseId(ctx *fiber.Ctx) error {
	courseId := ctx.Params("courseId")

	clos, err := c.courseLearningOutcomeUsecase.GetByCourseId(courseId)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusCreated, clos)
}

func (c courseLearningOutcomeController) Create(ctx *fiber.Ctx) error {
	var payload request.CreateCourseLearningOutcomePayload

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	err := c.courseLearningOutcomeUsecase.Create(payload.Code, payload.Description, payload.Weight, payload.SubProgramLearningOutcomeId, payload.ProgramOutcomeId, payload.CourseId, payload.Status)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusCreated, nil)
}

func (c courseLearningOutcomeController) Update(ctx *fiber.Ctx) error {
	var payload request.UpdateCourseLearningOutcomePayload

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	id := ctx.Params("cloId")

	err := c.courseLearningOutcomeUsecase.Update(id, &entity.CourseLearningOutcome{
		Code:                        payload.Code,
		Description:                 payload.Description,
		SubProgramLearningOutcomeId: payload.SubProgramLearningOutcomeId,
		ProgramOutcomeId:            payload.ProgramOutcomeId,
		Status:                      payload.Status,
	})

	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, nil)
}

func (c courseLearningOutcomeController) Delete(ctx *fiber.Ctx) error {
	cloId := ctx.Params("cloId")

	_, err := c.courseLearningOutcomeUsecase.GetById(cloId)
	if err != nil {
		return err
	}

	err = c.courseLearningOutcomeUsecase.Delete(cloId)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, nil)
}
