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

	faculty, err := c.FacultyUseCase.GetByID(facultyID)

	if err != nil {
		return err
	}

	return ctx.JSON(faculty)
}

func (c FacultyController) Create(ctx *fiber.Ctx) error {
	var faculty request.CreateFacultyRequestBody
	err := ctx.BodyParser(&faculty)

	if err != nil {
		return err
	}

	err, validationErrors := c.Validator.Validate(&faculty, ctx)
	if err != nil {
		return ctx.Status(400).JSON(validationErrors)
	}

	err = c.FacultyUseCase.Create(&entity.Faculty{
		Name: faculty.Name,
	})

	if err != nil {
		return err
	}

	return ctx.JSON(faculty)
}

func (c FacultyController) Update(ctx *fiber.Ctx) error {
	var faculty request.UpdateFacultyRequestBody
	err := ctx.BodyParser(&faculty)

	if err != nil {
		return err
	}

	err, validationErrors := c.Validator.Validate(&faculty, ctx)
	if err != nil {
		return ctx.Status(400).JSON(validationErrors)
	}

	err = c.FacultyUseCase.Update(&entity.Faculty{
		Name: faculty.Name,
	}, faculty.NewName)

	if err != nil {
		return err
	}

	return ctx.JSON(faculty)
}

func (c FacultyController) Delete(ctx *fiber.Ctx) error {
	var faculty request.DeleteFacultyRequestBody

	err := ctx.BodyParser(&faculty)
	if err != nil {
		return err
	}

	err, validationErrors := c.Validator.Validate(&faculty, ctx)
	if err != nil {
		return ctx.Status(400).JSON(validationErrors)
	}

	err = c.FacultyUseCase.Delete(faculty.Name)

	if err != nil {
		return err
	}

	return ctx.JSON(faculty.Name)
}
