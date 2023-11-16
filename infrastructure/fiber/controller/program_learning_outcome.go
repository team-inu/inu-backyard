package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team-inu/inu-backyard/entity"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/request"
	"github.com/team-inu/inu-backyard/internal/validator"
)

type programLearningOutcomeController struct {
	ProgramLearningOutcomeUsecase entity.ProgramLearningOutcomeUsecase
	Validator                     validator.PayloadValidator
}

func NewProgramLearningOutcomeController(programLearningOutcomeUsecase entity.ProgramLearningOutcomeUsecase) *programLearningOutcomeController {
	return &programLearningOutcomeController{
		ProgramLearningOutcomeUsecase: programLearningOutcomeUsecase,
		Validator:                     validator.NewPayloadValidator(),
	}
}

func (c programLearningOutcomeController) GetAll(ctx *fiber.Ctx) error {
	plos, err := c.ProgramLearningOutcomeUsecase.GetAll()
	if err != nil {
		return err
	}

	return ctx.JSON(plos)
}

func (c programLearningOutcomeController) GetByID(ctx *fiber.Ctx) error {
	ploId := ctx.Params("ploId")

	plos, err := c.ProgramLearningOutcomeUsecase.GetByID(ploId)
	if err != nil {
		return err
	}

	return ctx.JSON(plos)
}

func (c programLearningOutcomeController) Create(ctx *fiber.Ctx) error {
	var plo request.CreateProgramLearningOutcomeBody
	err := ctx.BodyParser(&plo)
	if err != nil {
		return err
	}

	err, validationErrors := c.Validator.Validate(plo, ctx)
	if err != nil {
		return ctx.JSON(validationErrors)
	}

	createdClo, err := c.ProgramLearningOutcomeUsecase.Create(plo.Code, plo.DescriptionThai, plo.DescriptionEng, plo.ProgramYear)
	if err != nil {
		return err
	}

	return ctx.JSON(createdClo)
}

func (c programLearningOutcomeController) Delete(ctx *fiber.Ctx) error {
	ploId := ctx.Params("ploId")

	_, err := c.ProgramLearningOutcomeUsecase.GetByID(ploId)
	if err != nil {
		return err
	}

	err = c.ProgramLearningOutcomeUsecase.Delete(ploId)
	if err != nil {
		return err
	}

	return ctx.SendStatus(200)
}
