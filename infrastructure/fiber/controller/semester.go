package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team-inu/inu-backyard/entity"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/request"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/response"
	"github.com/team-inu/inu-backyard/internal/validator"
)

type semesterController struct {
	SemesterUseCase entity.SemesterUseCase
	Validator       validator.PayloadValidator
}

func NewSemesterController(semesterUseCase entity.SemesterUseCase) *semesterController {
	return &semesterController{
		SemesterUseCase: semesterUseCase,
		Validator:       validator.NewPayloadValidator(),
	}
}

func (c semesterController) GetAll(ctx *fiber.Ctx) error {
	semesters, err := c.SemesterUseCase.GetAll()
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, semesters)
}

func (c semesterController) GetByID(ctx *fiber.Ctx) error {
	semesterID := ctx.Params("semesterID")

	semester, err := c.SemesterUseCase.GetByID(semesterID)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, semester)
}

func (c semesterController) Create(ctx *fiber.Ctx) error {
	var payload request.CreateSemesterPayload

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	err := c.SemesterUseCase.Create(payload.Year, payload.SemesterSequence)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusCreated, nil)
}

func (c semesterController) Update(ctx *fiber.Ctx) error {
	semesterID := ctx.Params("semesterID")
	var payload request.UpdateSemesterPayload

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	err := c.SemesterUseCase.Update(&entity.Semester{
		ID:               semesterID,
		Year:             payload.Year,
		SemesterSequence: payload.SemesterSequence,
	})
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, nil)
}

func (c semesterController) Delete(ctx *fiber.Ctx) error {
	semesterID := ctx.Params("semesterID")

	err := c.SemesterUseCase.Delete(semesterID)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, nil)
}
