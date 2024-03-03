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

func NewSemesterController(validator validator.PayloadValidator, semesterUseCase entity.SemesterUseCase) *semesterController {
	return &semesterController{
		SemesterUseCase: semesterUseCase,
		Validator:       validator,
	}
}

func (c semesterController) GetAll(ctx *fiber.Ctx) error {
	semesters, err := c.SemesterUseCase.GetAll()
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, semesters)
}

func (c semesterController) GetById(ctx *fiber.Ctx) error {
	semesterId := ctx.Params("semesterId")

	semester, err := c.SemesterUseCase.GetById(semesterId)
	if err != nil {
		return err
	}

	if semester == nil {
		return response.NewSuccessResponse(ctx, fiber.StatusNotFound, semester)
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
	semesterId := ctx.Params("semesterId")
	var payload request.UpdateSemesterPayload

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	err := c.SemesterUseCase.Update(&entity.Semester{
		Id:               semesterId,
		Year:             payload.Year,
		SemesterSequence: payload.SemesterSequence,
	})
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, nil)
}

func (c semesterController) Delete(ctx *fiber.Ctx) error {
	semesterId := ctx.Params("semesterId")

	err := c.SemesterUseCase.Delete(semesterId)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, nil)
}
