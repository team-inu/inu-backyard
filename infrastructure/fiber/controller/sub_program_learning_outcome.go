package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team-inu/inu-backyard/entity"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/request"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/response"
	"github.com/team-inu/inu-backyard/internal/validator"
)

type subProgramLearningOutcomeController struct {
	programLearningOutcomeUseCase entity.ProgramLearningOutcomeUseCase
	Validator                     validator.PayloadValidator
}

func NewSubProgramLearningOutcomeController(validator validator.PayloadValidator, programLearningOutcomeUseCase entity.ProgramLearningOutcomeUseCase) *subProgramLearningOutcomeController {
	return &subProgramLearningOutcomeController{
		programLearningOutcomeUseCase: programLearningOutcomeUseCase,
		Validator:                     validator,
	}
}

func (c subProgramLearningOutcomeController) GetAll(ctx *fiber.Ctx) error {
	splos, err := c.programLearningOutcomeUseCase.GetAllSubPlo()
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, splos)
}

func (c subProgramLearningOutcomeController) GetById(ctx *fiber.Ctx) error {
	sploId := ctx.Params("sploId")

	splo, err := c.programLearningOutcomeUseCase.GetSubPLO(sploId)
	if err != nil {
		return err
	}

	if splo == nil {
		return response.NewSuccessResponse(ctx, fiber.StatusNotFound, splo)
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, splo)
}

func (c subProgramLearningOutcomeController) Create(ctx *fiber.Ctx) error {
	var payload request.CreateSubProgramLearningOutcomePayload
	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	subPlos := make([]entity.CreateSubProgramLearningOutcomeDto, 0, len(payload.SubProgramLearningOutcomes))
	for _, subPlo := range payload.SubProgramLearningOutcomes {
		subPlos = append(subPlos, entity.CreateSubProgramLearningOutcomeDto(subPlo))
	}

	err := c.programLearningOutcomeUseCase.CreateSubPLO(subPlos)
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

	id := ctx.Params("sploId")

	err := c.programLearningOutcomeUseCase.UpdateSubPLO(id, &entity.SubProgramLearningOutcome{
		Code:                     payload.SubProgramLearningOutcomes[0].Code,
		DescriptionThai:          payload.SubProgramLearningOutcomes[0].DescriptionThai,
		DescriptionEng:           *payload.SubProgramLearningOutcomes[0].DescriptionEng,
		ProgramLearningOutcomeId: payload.SubProgramLearningOutcomes[0].ProgramLearningOutcomeId,
	})

	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, nil)
}

func (c subProgramLearningOutcomeController) Delete(ctx *fiber.Ctx) error {
	sploId := ctx.Params("sploId")

	err := c.programLearningOutcomeUseCase.DeleteSubPLO(sploId)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, nil)
}
