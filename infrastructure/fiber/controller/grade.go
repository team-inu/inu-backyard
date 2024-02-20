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

	return ctx.JSON(grades)
}

func (c gradeController) GetByID(ctx *fiber.Ctx) error {
	gradeID := ctx.Params("gradeID")

	grade, err := c.gradeUseCase.GetByID(gradeID)

	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusCreated, grade)
}

func (c gradeController) Create(ctx *fiber.Ctx) error {
	var payload request.CreateGradePayload

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	err := c.gradeUseCase.Create(payload.StudentID, payload.Year, payload.Grade)

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

	id := ctx.Params("gradeID")

	err := c.gradeUseCase.Update(id, &entity.Grade{
		StudentID: payload.StudentID,
		Grade:     payload.Grade,
	})

	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, nil)
}

func (c gradeController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("gradeID")

	err := c.gradeUseCase.Delete(id)

	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, nil)
}
