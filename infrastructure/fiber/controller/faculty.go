package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team-inu/inu-backyard/entity"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/request"
	"github.com/team-inu/inu-backyard/internal/validator"
)

type FacultyController struct {
	FacultyUseCase entity.FacultyUseCase
	Validator      validator.PayloadValidator
}

func NewFacultyController(facultyUseCase entity.FacultyUseCase) *FacultyController {
	return &FacultyController{
		FacultyUseCase: facultyUseCase,
		Validator:      validator.NewPayloadValidator(),
	}
}

func (c FacultyController) GetAll(ctx *fiber.Ctx) error {
	faculties, err := c.FacultyUseCase.GetAll()
	if err != nil {
		return err
	}

	return ctx.JSON(faculties)
}

func (c FacultyController) GetByID(ctx *fiber.Ctx) error {
	facultyID := ctx.Params("facultyID")

	faculty, err := c.FacultyUseCase.GetByName(facultyID)

	if err != nil {
		return err
	}

	return ctx.JSON(faculty)
}

func (c FacultyController) Create(ctx *fiber.Ctx) error {
	var payload request.CreateFacultyRequestPayload

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	err := c.FacultyUseCase.Create(&entity.Faculty{
		Name: payload.Name,
	})

	if err != nil {
		return err
	}

	return ctx.JSON(payload)
}

func (c FacultyController) Update(ctx *fiber.Ctx) error {
	var payload request.UpdateFacultyRequestPayload

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	err := c.FacultyUseCase.Update(&entity.Faculty{
		Name: payload.Name,
	}, payload.NewName)

	if err != nil {
		return err
	}

	return ctx.JSON(payload)
}

func (c FacultyController) Delete(ctx *fiber.Ctx) error {
	var payload request.DeleteFacultyRequestPayload

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	err := c.FacultyUseCase.Delete(payload.Name)

	if err != nil {
		return err
	}

	return ctx.JSON(payload.Name)
}
