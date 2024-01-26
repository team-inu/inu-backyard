package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team-inu/inu-backyard/entity"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/request"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/response"
	"github.com/team-inu/inu-backyard/internal/validator"
)

type programLearningOutcomeController struct {
	programLearningOutcomeUsecase entity.ProgramLearningOutcomeUsecase
	Validator                     validator.PayloadValidator
}

func NewProgramLearningOutcomeController(validator validator.PayloadValidator, programLearningOutcomeUsecase entity.ProgramLearningOutcomeUsecase) *programLearningOutcomeController {
	return &programLearningOutcomeController{
		programLearningOutcomeUsecase: programLearningOutcomeUsecase,
		Validator:                     validator,
	}
}

func (c programLearningOutcomeController) GetAll(ctx *fiber.Ctx) error {
	plos, err := c.programLearningOutcomeUsecase.GetAll()
	if err != nil {
		return err
	}

	return ctx.JSON(plos)
}

func (c programLearningOutcomeController) GetByID(ctx *fiber.Ctx) error {
	ploId := ctx.Params("ploId")

	plo, err := c.programLearningOutcomeUsecase.GetByID(ploId)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusCreated, plo)
}

func (c programLearningOutcomeController) Create(ctx *fiber.Ctx) error {
	var payload request.CreateProgramLearningOutcomePayload

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	err := c.programLearningOutcomeUsecase.Create(payload.Code, payload.DescriptionThai, payload.DescriptionEng, payload.ProgramYear)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusCreated, nil)
}

func (c programLearningOutcomeController) Update(ctx *fiber.Ctx) error {
	var payload request.UpdateProgramLearningOutcomePayload

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	id := ctx.Params("ploID")

	err := c.programLearningOutcomeUsecase.Update(id, &entity.ProgramLearningOutcome{
		Code:            payload.Code,
		DescriptionThai: payload.DescriptionThai,
		DescriptionEng:  payload.DescriptionEng,
		ProgramYear:     payload.ProgramYear,
	})

	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, nil)
}

func (c programLearningOutcomeController) Delete(ctx *fiber.Ctx) error {
	ploId := ctx.Params("ploId")

	_, err := c.programLearningOutcomeUsecase.GetByID(ploId)
	if err != nil {
		return err
	}

	err = c.programLearningOutcomeUsecase.Delete(ploId)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, nil)
}
