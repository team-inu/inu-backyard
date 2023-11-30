package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team-inu/inu-backyard/entity"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/request"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/response"
	"github.com/team-inu/inu-backyard/internal/validator"
)

type subProgramLearningOutcomeController struct {
	subProgramLearningOutcomeUsecase entity.SubProgramLearningOutcomeUsecase
	Validator                        validator.PayloadValidator
}

func NewSubProgramLearningOutcomeController(subProgramLearningOutcomeUsecase entity.SubProgramLearningOutcomeUsecase) *subProgramLearningOutcomeController {
	return &subProgramLearningOutcomeController{
		subProgramLearningOutcomeUsecase: subProgramLearningOutcomeUsecase,
		Validator:                        validator.NewPayloadValidator(),
	}
}

func (c subProgramLearningOutcomeController) GetAll(ctx *fiber.Ctx) error {
	splos, err := c.subProgramLearningOutcomeUsecase.GetAll()
	if err != nil {
		return err
	}

	return ctx.JSON(splos)
}

func (c subProgramLearningOutcomeController) GetByID(ctx *fiber.Ctx) error {
	sploId := ctx.Params("sploId")

	splo, err := c.subProgramLearningOutcomeUsecase.GetByID(sploId)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusCreated, splo)
}

func (c subProgramLearningOutcomeController) Create(ctx *fiber.Ctx) error {
	var payload request.CreateSubProgramLearningOutcomePayload

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	err := c.subProgramLearningOutcomeUsecase.Create(payload.Code, payload.DescriptionThai, payload.DescriptionEng, payload.ProgramLearningOutcomeID)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusCreated, nil)
}

func (c subProgramLearningOutcomeController) Update(ctx *fiber.Ctx) error {
	var payload request.UpdateSubProgramLearningOutcomePayload

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	id := ctx.Params("sploID")

	err := c.subProgramLearningOutcomeUsecase.Update(id, &entity.SubProgramLearningOutcome{
		Code:                     payload.Code,
		DescriptionThai:          payload.DescriptionThai,
		DescriptionEng:           payload.DescriptionEng,
		ProgramLearningOutcomeID: payload.ProgramLearningOutcomeID,
	})

	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, nil)
}

func (c subProgramLearningOutcomeController) Delete(ctx *fiber.Ctx) error {
	sploId := ctx.Params("sploId")

	_, err := c.subProgramLearningOutcomeUsecase.GetByID(sploId)
	if err != nil {
		return err
	}

	err = c.subProgramLearningOutcomeUsecase.Delete(sploId)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, nil)
}
