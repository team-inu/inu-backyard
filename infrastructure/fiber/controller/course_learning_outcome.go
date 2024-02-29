package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team-inu/inu-backyard/entity"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/request"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/response"
	"github.com/team-inu/inu-backyard/internal/validator"
)

type courseLearningOutcomeController struct {
	courseLearningOutcomeUseCase entity.CourseLearningOutcomeUseCase
	Validator                    validator.PayloadValidator
}

func NewCourseLearningOutcomeController(validator validator.PayloadValidator, courseLearningOutcomeUseCase entity.CourseLearningOutcomeUseCase) *courseLearningOutcomeController {
	return &courseLearningOutcomeController{
		courseLearningOutcomeUseCase: courseLearningOutcomeUseCase,
		Validator:                    validator,
	}
}

func (c courseLearningOutcomeController) GetAll(ctx *fiber.Ctx) error {
	clos, err := c.courseLearningOutcomeUseCase.GetAll()
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusCreated, clos)

}

func (c courseLearningOutcomeController) GetById(ctx *fiber.Ctx) error {
	cloId := ctx.Params("cloId")

	clo, err := c.courseLearningOutcomeUseCase.GetById(cloId)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusCreated, clo)
}

func (c courseLearningOutcomeController) GetByCourseId(ctx *fiber.Ctx) error {
	courseId := ctx.Params("courseId")

	clos, err := c.courseLearningOutcomeUseCase.GetByCourseId(courseId)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusCreated, clos)
}

func (c courseLearningOutcomeController) Create(ctx *fiber.Ctx) error {
	var payload request.CreateCourseLearningOutcomePayload
	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	err := c.courseLearningOutcomeUseCase.Create(entity.CreateCourseLearningOutcomeDto{
		Code:                                payload.Code,
		Description:                         payload.Description,
		Status:                              payload.Status,
		ExpectedPassingAssignmentPercentage: payload.ExpectedPassingAssignmentPercentage,
		ExpectedScorePercentage:             payload.ExpectedScorePercentage,
		ExpectedPassingStudentPercentage:    payload.ExpectedPassingStudentPercentage,
		CourseId:                            payload.CourseId,
		ProgramOutcomeId:                    payload.ProgramOutcomeId,
		SubProgramLearningOutcomeIds:        payload.SubProgramLearningOutcomeIds,
	})
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusCreated, nil)
}

func (c courseLearningOutcomeController) Update(ctx *fiber.Ctx) error {
	var payload request.UpdateCourseLearningOutcomePayload

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	id := ctx.Params("cloId")

	err := c.courseLearningOutcomeUseCase.Update(id, entity.UpdateCourseLeaningOutcomeDto{
		Code:                                payload.Code,
		Description:                         payload.Description,
		ExpectedPassingAssignmentPercentage: payload.ExpectedPassingAssignmentPercentage,
		ExpectedScorePercentage:             payload.ExpectedScorePercentage,
		ExpectedPassingStudentPercentage:    payload.ExpectedPassingStudentPercentage,
		Status:                              payload.Status,
		ProgramOutcomeId:                    payload.ProgramOutcomeId,
	})

	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, nil)
}

func (c courseLearningOutcomeController) Delete(ctx *fiber.Ctx) error {
	cloId := ctx.Params("cloId")

	_, err := c.courseLearningOutcomeUseCase.GetById(cloId)
	if err != nil {
		return err
	}

	err = c.courseLearningOutcomeUseCase.Delete(cloId)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, nil)
}
