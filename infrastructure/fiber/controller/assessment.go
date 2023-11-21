package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team-inu/inu-backyard/entity"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/request"
	"github.com/team-inu/inu-backyard/internal/validator"
)

type assessmentController struct {
	AssessmentUseCase entity.AssessmentUseCase
	Validator         validator.PayloadValidator
}

func NewAssessmentController(assessmentUseCase entity.AssessmentUseCase) *assessmentController {
	return &assessmentController{
		AssessmentUseCase: assessmentUseCase,
		Validator:         validator.NewPayloadValidator(),
	}
}

func (c assessmentController) GetByID(ctx *fiber.Ctx) error {
	assessmentID := ctx.Params("assessmentID")

	assessment, err := c.AssessmentUseCase.GetByID(assessmentID)

	if err != nil {
		return err
	}

	return ctx.JSON(assessment)
}

func (c assessmentController) GetAssessments(ctx *fiber.Ctx) error {
	var payload request.GetAssessmentsByParamsPayload

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	assessment, err := c.AssessmentUseCase.GetByParams(&entity.Assessment{
		CourseLearningOutcomeID: payload.CourseLearningOutcomeID,
	}, -1, -1)

	if err != nil {
		return err
	}

	return ctx.JSON(assessment)
}

func (c assessmentController) GetAssessmentsByCourseID(ctx *fiber.Ctx) error {
	var payload request.GetAssessmentsByCourseIDPayload

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	assessment, err := c.AssessmentUseCase.GetByCourseID(payload.CourseID, -1, -1)

	if err != nil {
		return err
	}

	return ctx.JSON(assessment)
}

func (c assessmentController) Create(ctx *fiber.Ctx) error {
	var payload request.CreateAssessmentPayload
	err := ctx.BodyParser(&payload)
	if err != nil {
		return err
	}

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	err = c.AssessmentUseCase.Create(&entity.Assessment{
		Name:                    payload.Name,
		Description:             payload.Description,
		Score:                   *payload.Score,
		Weight:                  *payload.Weight,
		CourseLearningOutcomeID: payload.CourseLearningOutcomeID,
	})
	if err != nil {
		return err
	}

	return ctx.JSON(payload)
}

func (c assessmentController) CreateMany(ctx *fiber.Ctx) error {
	var payload request.CreateBulkAssessmentsPayload
	err := ctx.BodyParser(&payload)

	if err != nil {
		return err
	}

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	newAssessments := []entity.Assessment{}

	for _, assessment := range payload.Assessments {
		newAssessments = append(newAssessments, entity.Assessment{
			Name:                    assessment.Name,
			Description:             assessment.Description,
			Score:                   *assessment.Score,
			Weight:                  *assessment.Weight,
			CourseLearningOutcomeID: assessment.CourseLearningOutcomeID,
		})
	}

	err = c.AssessmentUseCase.CreateMany(newAssessments)
	if err != nil {
		return err
	}

	return ctx.JSON(payload)
}

func (c assessmentController) Update(ctx *fiber.Ctx) error {
	var payload request.UpdateAssessmentRequestPayload

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	err := c.AssessmentUseCase.Update(payload.ID, &entity.Assessment{
		Name:                    payload.Name,
		Description:             payload.Description,
		Score:                   payload.Score,
		Weight:                  payload.Weight,
		CourseLearningOutcomeID: payload.CourseLearningOutcomeID,
	})

	if err != nil {
		return err
	}

	return ctx.JSON(payload)
}

func (c assessmentController) Delete(ctx *fiber.Ctx) error {
	var payload request.DeleteAssessmentRequestPayload

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	err := c.AssessmentUseCase.Delete(payload.ID)

	if err != nil {
		return err
	}

	return ctx.JSON(payload.ID)
}
