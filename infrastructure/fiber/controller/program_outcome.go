package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team-inu/inu-backyard/entity"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/request"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/response"
	"github.com/team-inu/inu-backyard/internal/validator"
)

type programOutcomeController struct {
	programOutcomeUseCase entity.ProgramOutcomeUseCase
	Validator             validator.PayloadValidator
}

func NewProgramOutcomeController(validator validator.PayloadValidator, programOutcomeUseCase entity.ProgramOutcomeUseCase) *programOutcomeController {
	return &programOutcomeController{
		programOutcomeUseCase: programOutcomeUseCase,
		Validator:             validator,
	}
}

func (c programOutcomeController) GetAll(ctx *fiber.Ctx) error {
	pos, err := c.programOutcomeUseCase.GetAll()
	if err != nil {
		return err
	}

	return ctx.JSON(pos)
}

func (c programOutcomeController) GetById(ctx *fiber.Ctx) error {
	poId := ctx.Params("poId")

	po, err := c.programOutcomeUseCase.GetById(poId)
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

	pos := make([]entity.ProgramOutcome, 0, len(payload.ProgramOutcomes))
	for _, po := range payload.ProgramOutcomes {

		pos = append(pos, entity.ProgramOutcome{
			Code:        po.Code,
			Name:        po.Name,
			Description: po.Description,
		})
	}

	err := c.programOutcomeUseCase.Create(pos)
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

	id := ctx.Params("poId")

	err := c.programOutcomeUseCase.Update(id, &entity.ProgramOutcome{
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

	_, err := c.programOutcomeUseCase.GetById(poId)
	if err != nil {
		return err
	}

	err = c.programOutcomeUseCase.Delete(poId)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, nil)
}
