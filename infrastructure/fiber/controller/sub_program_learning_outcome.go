package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team-inu/inu-backyard/entity"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/request"
	"github.com/team-inu/inu-backyard/internal/validator"
)

type subProgramLearningOutcomeController struct {
	SubProgramLearningOutcomeUsecase entity.SubProgramLearningOutcomeUsecase
	Validator                        validator.PayloadValidator
}

func NewSubProgramLearningOutcomeController(subProgramLearningOutcomeUsecase entity.SubProgramLearningOutcomeUsecase) *subProgramLearningOutcomeController {
	return &subProgramLearningOutcomeController{
		SubProgramLearningOutcomeUsecase: subProgramLearningOutcomeUsecase,
		Validator:                        validator.NewPayloadValidator(),
	}
}

func (c subProgramLearningOutcomeController) GetAll(ctx *fiber.Ctx) error {
	splos, err := c.SubProgramLearningOutcomeUsecase.GetAll()
	if err != nil {
		return err
	}

	return ctx.JSON(splos)
}

func (c subProgramLearningOutcomeController) GetByID(ctx *fiber.Ctx) error {
	sploId := ctx.Params("sploId")

	splos, err := c.SubProgramLearningOutcomeUsecase.GetByID(sploId)
	if err != nil {
		return err
	}

	return ctx.JSON(splos)
}

func (c subProgramLearningOutcomeController) Create(ctx *fiber.Ctx) error {
	var payload request.CreateSubProgramLearningOutcomeBody
	err := ctx.BodyParser(&payload)
	if err != nil {
		return err
	}

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	createdClo, err := c.SubProgramLearningOutcomeUsecase.Create(payload.Code, payload.DescriptionThai, payload.DescriptionEng, payload.ProgramLearningOutcomeID)
	if err != nil {
		return err
	}

	return ctx.JSON(createdClo)
}

func (c subProgramLearningOutcomeController) Delete(ctx *fiber.Ctx) error {
	sploId := ctx.Params("sploId")

	_, err := c.SubProgramLearningOutcomeUsecase.GetByID(sploId)
	if err != nil {
		return err
	}

	err = c.SubProgramLearningOutcomeUsecase.Delete(sploId)
	if err != nil {
		return err
	}

	return ctx.SendStatus(200)
}
