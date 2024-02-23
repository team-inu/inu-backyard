package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team-inu/inu-backyard/entity"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/request"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/response"
	"github.com/team-inu/inu-backyard/internal/validator"
)

type programmeController struct {
	programmeUseCase entity.ProgrammeUseCase
	Validator        validator.PayloadValidator
}

func NewProgrammeController(validator validator.PayloadValidator, programmeUseCase entity.ProgrammeUseCase) *programmeController {
	return &programmeController{
		programmeUseCase: programmeUseCase,
		Validator:        validator,
	}
}

func (c programmeController) GetAll(ctx *fiber.Ctx) error {
	programmes, err := c.programmeUseCase.GetAll()
	if err != nil {
		return err
	}

	return ctx.JSON(programmes)
}

func (c programmeController) GetByName(ctx *fiber.Ctx) error {
	name := ctx.Params("programmeName")

	programme, err := c.programmeUseCase.Get(name)

	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusCreated, programme)
}

func (c programmeController) Create(ctx *fiber.Ctx) error {
	var payload request.CreateProgrammePayload

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	err := c.programmeUseCase.Create(payload.Name)

	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusCreated, nil)
}

func (c programmeController) Update(ctx *fiber.Ctx) error {
	var payload request.UpdateProgrammePayload

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	name := ctx.Params("programmeName")

	err := c.programmeUseCase.Update(name, &entity.Programme{
		Name: payload.Name,
	})

	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, nil)
}

func (c programmeController) Delete(ctx *fiber.Ctx) error {
	name := ctx.Params("programmeName")

	err := c.programmeUseCase.Delete(name)

	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, nil)
}
