package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team-inu/inu-backyard/entity"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/request"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/response"
	"github.com/team-inu/inu-backyard/internal/validator"
)

type assignmentController struct {
	AssignmentUseCase entity.AssignmentUseCase
	Validator         validator.PayloadValidator
}

func NewAssignmentController(validator validator.PayloadValidator, assignmentUseCase entity.AssignmentUseCase) *assignmentController {
	return &assignmentController{
		AssignmentUseCase: assignmentUseCase,
		Validator:         validator,
	}
}

func (c assignmentController) GetById(ctx *fiber.Ctx) error {
	assignmentId := ctx.Params("assignmentId")

	assignment, err := c.AssignmentUseCase.GetById(assignmentId)

	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, assignment)
}

func (c assignmentController) GetAssignments(ctx *fiber.Ctx) error {
	var payload request.GetAssignmentsByParamsPayload

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	assignments, err := c.AssignmentUseCase.GetByParams(&entity.Assignment{
		// CourseLearningOutcomeId: payload.CourseLearningOutcomeId,
	}, -1, -1)

	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, assignments)
}

func (c assignmentController) GetAssignmentsByCourseId(ctx *fiber.Ctx) error {
	var payload request.GetAssignmentsByCourseIdPayload

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	assignment, err := c.AssignmentUseCase.GetByCourseId(payload.CourseId, -1, -1)

	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, assignment)
}

func (c assignmentController) Create(ctx *fiber.Ctx) error {
	var payload request.CreateAssignmentPayload

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	err := c.AssignmentUseCase.Create(
		payload.Name,
		payload.Description,
		*payload.MaxScore,
		*payload.Weight,
		*payload.ExpectedScorePercentage,
		*payload.ExpectedPassingStudentPercentage,
		payload.CourseLearningOutcomeIds,
	)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusCreated, nil)
}

func (c assignmentController) CreateMany(ctx *fiber.Ctx) error {
	var payload request.CreateBulkAssignmentsPayload

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	newAssignments := []entity.Assignment{}

	for _, assignment := range payload.Assignments {
		newAssignments = append(newAssignments, entity.Assignment{
			Name:        assignment.Name,
			Description: assignment.Description,
			Weight:      *assignment.Weight,
			// CourseLearningOutcomeId: assignment.CourseLearningOutcomeId,
		})
	}

	err := c.AssignmentUseCase.CreateMany(newAssignments)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusCreated, nil)
}

func (c assignmentController) Update(ctx *fiber.Ctx) error {
	var payload request.UpdateAssignmentRequestPayload

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	err := c.AssignmentUseCase.Update(payload.Id, &entity.Assignment{
		Name:        payload.Name,
		Description: payload.Description,
		Weight:      payload.Weight,
		// CourseLearningOutcomeId: payload.CourseLearningOutcomeId,
	})

	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, nil)
}

func (c assignmentController) Delete(ctx *fiber.Ctx) error {
	var payload request.DeleteAssignmentRequestPayload

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	err := c.AssignmentUseCase.Delete(payload.Id)

	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, nil)
}
