package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team-inu/inu-backyard/entity"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/request"
	"github.com/team-inu/inu-backyard/internal/validator"
)

type programOutcomeController struct {
	ProgramOutcomeUsecase entity.ProgramOutcomeUsecase
	Validator             validator.PayloadValidator
}

func NewProgramOutcomeController(programOutcomeUsecase entity.ProgramOutcomeUsecase) *programOutcomeController {
	return &programOutcomeController{
		ProgramOutcomeUsecase: programOutcomeUsecase,
		Validator:             validator.NewPayloadValidator(),
	}
}

func (c programOutcomeController) GetAll(ctx *fiber.Ctx) error {
	pos, err := c.ProgramOutcomeUsecase.GetAll()
	if err != nil {
		return err
	}

	return ctx.JSON(pos)
}

func (c programOutcomeController) GetByID(ctx *fiber.Ctx) error {
	poId := ctx.Params("poId")

	pos, err := c.ProgramOutcomeUsecase.GetByID(poId)
	if err != nil {
		return err
	}

	return ctx.JSON(pos)
}

func (c programOutcomeController) Create(ctx *fiber.Ctx) error {
	var payload request.CreateProgramOutcomePayload
	err := ctx.BodyParser(&payload)
	if err != nil {
		return err
	}

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	createdClo, err := c.ProgramOutcomeUsecase.Create(payload.Code, payload.Name, payload.Description)
	if err != nil {
		return err
	}

	return ctx.JSON(createdClo)
}

func (c programOutcomeController) Delete(ctx *fiber.Ctx) error {
	poId := ctx.Params("poId")

	_, err := c.ProgramOutcomeUsecase.GetByID(poId)
	if err != nil {
		return err
	}

	err = c.ProgramOutcomeUsecase.Delete(poId)
	if err != nil {
		return err
	}

	return ctx.SendStatus(200)
}
