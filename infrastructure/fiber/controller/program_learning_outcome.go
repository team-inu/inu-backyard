package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team-inu/inu-backyard/entity"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/request"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/response"
	"github.com/team-inu/inu-backyard/internal/validator"
)

type programLearningOutcomeController struct {
	programLearningOutcomeUseCase entity.ProgramLearningOutcomeUseCase
	Validator                     validator.PayloadValidator
}

func NewProgramLearningOutcomeController(validator validator.PayloadValidator, programLearningOutcomeUseCase entity.ProgramLearningOutcomeUseCase) *programLearningOutcomeController {
	return &programLearningOutcomeController{
		programLearningOutcomeUseCase: programLearningOutcomeUseCase,
		Validator:                     validator,
	}
}

func (c programLearningOutcomeController) GetAll(ctx *fiber.Ctx) error {
	plos, err := c.programLearningOutcomeUseCase.GetAll()
	if err != nil {
		return err
	}

	return ctx.JSON(plos)
}

func (c programLearningOutcomeController) GetById(ctx *fiber.Ctx) error {
	ploId := ctx.Params("ploId")

	plo, err := c.programLearningOutcomeUseCase.GetById(ploId)
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

	plos := []entity.CrateProgramLearningOutcomeDto{}
	for _, plo := range payload.ProgramLearningOutcomes {
		plos = append(plos, entity.CrateProgramLearningOutcomeDto{
			Code:            plo.Code,
			DescriptionThai: plo.DescriptionThai,
			DescriptionEng:  plo.DescriptionEng,
			ProgramYear:     plo.ProgramYear,
			ProgrammeName:   plo.ProgrammeName,
		})
	}

	err := c.programLearningOutcomeUseCase.Create(plos)
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

	id := ctx.Params("ploId")

	err := c.programLearningOutcomeUseCase.Update(id, &entity.ProgramLearningOutcome{
		Code:            payload.Code,
		DescriptionThai: payload.DescriptionThai,
		DescriptionEng:  payload.DescriptionEng,
		ProgramYear:     payload.ProgramYear,
		ProgrammeId:     payload.Programme,
	})

	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, nil)
}

func (c programLearningOutcomeController) Delete(ctx *fiber.Ctx) error {
	ploId := ctx.Params("ploId")

	_, err := c.programLearningOutcomeUseCase.GetById(ploId)
	if err != nil {
		return err
	}

	err = c.programLearningOutcomeUseCase.Delete(ploId)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, nil)
}
