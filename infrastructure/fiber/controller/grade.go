package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team-inu/inu-backyard/entity"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/request"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/response"
	"github.com/team-inu/inu-backyard/internal/validator"
)

type gradeController struct {
	gradeUseCase entity.GradeUseCase
	Validator    validator.PayloadValidator
}

func NewGradeController(validator validator.PayloadValidator, gradeUseCase entity.GradeUseCase) *gradeController {
	return &gradeController{
		gradeUseCase: gradeUseCase,
		Validator:    validator,
	}
}

func (c gradeController) GetAll(ctx *fiber.Ctx) error {
	grades, err := c.gradeUseCase.GetAll()
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, grades)
}

func (c gradeController) GetById(ctx *fiber.Ctx) error {
	gradeId := ctx.Params("gradeId")

	grade, err := c.gradeUseCase.GetById(gradeId)

	if err != nil {
		return err
	}

	if grade == nil {
		return response.NewSuccessResponse(ctx, fiber.StatusNotFound, grade)
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, grade)
}

func (c gradeController) Create(ctx *fiber.Ctx) error {
	var payload request.CreateGradePayload

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	err := c.gradeUseCase.Create(payload.StudentId, payload.Year, payload.Grade)

	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusCreated, nil)
}

func (c gradeController) Update(ctx *fiber.Ctx) error {
	var payload request.UpdateGradePayload

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	id := ctx.Params("gradeId")

	err := c.gradeUseCase.Update(id, &entity.Grade{
		StudentId: payload.StudentId,
		Grade:     payload.Grade,
	})

	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, nil)
}

func (c gradeController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("gradeId")

	err := c.gradeUseCase.Delete(id)

	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, nil)
}
