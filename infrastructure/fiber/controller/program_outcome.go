package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team-inu/inu-backyard/entity"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/request"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/response"
	"github.com/team-inu/inu-backyard/internal/validator"
)

type programOutcomeController struct {
	programOutcomeUsecase entity.ProgramOutcomeUsecase
	Validator             validator.PayloadValidator
}

func NewProgramOutcomeController(programOutcomeUsecase entity.ProgramOutcomeUsecase) *programOutcomeController {
	return &programOutcomeController{
		programOutcomeUsecase: programOutcomeUsecase,
		Validator:             validator.NewPayloadValidator(),
	}
}

func (c programOutcomeController) GetAll(ctx *fiber.Ctx) error {
	pos, err := c.programOutcomeUsecase.GetAll()
	if err != nil {
		return err
	}

	return ctx.JSON(pos)
}

func (c programOutcomeController) GetByID(ctx *fiber.Ctx) error {
	poId := ctx.Params("poId")

	po, err := c.programOutcomeUsecase.GetByID(poId)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusCreated, po)
}

func (c programOutcomeController) Create(ctx *fiber.Ctx) error {
	var payload request.CreateProgramOutcomePayload

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	err := c.programOutcomeUsecase.Create(payload.Code, payload.Name, payload.Description)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusCreated, nil)
}

func (c programOutcomeController) Update(ctx *fiber.Ctx) error {
	var payload request.UpdateProgramOutcomePayload

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	id := ctx.Params("poID")

	err := c.programOutcomeUsecase.Update(id, &entity.ProgramOutcome{
		Code:        payload.Code,
		Name:        payload.Name,
		Description: payload.Description,
	})

	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, nil)
}

func (c programOutcomeController) Delete(ctx *fiber.Ctx) error {
	poId := ctx.Params("poId")

	_, err := c.programOutcomeUsecase.GetByID(poId)
	if err != nil {
		return err
	}

	err = c.programOutcomeUsecase.Delete(poId)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, nil)
}
