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

	if assignment == nil {
		return response.NewSuccessResponse(ctx, fiber.StatusNotFound, assignment)
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

func (c assignmentController) GetByCourseId(ctx *fiber.Ctx) error {
	var payload request.GetAssignmentsByCourseIdPayload

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	assignments, err := c.AssignmentUseCase.GetByCourseId(payload.CourseId)

	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, assignments)
}

func (c assignmentController) Create(ctx *fiber.Ctx) error {
	var payload request.CreateAssignmentPayload

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	err := c.AssignmentUseCase.Create(
		payload.AssignmentGroupId,
		payload.Name,
		payload.Description,
		*payload.MaxScore,
		*payload.Weight,
		*payload.ExpectedScorePercentage,
		*payload.ExpectedPassingStudentPercentage,
		payload.CourseLearningOutcomeIds,
		*payload.IsIncludedInClo,
	)
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

	id := ctx.Params("assignmentId")

	err := c.AssignmentUseCase.Update(id, payload.Name, payload.Description, *payload.MaxScore, *payload.Weight, *payload.ExpectedScorePercentage, *payload.ExpectedPassingStudentPercentage, *payload.IsIncludedInClo)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, nil)
}

func (c assignmentController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("assignmentId")

	err := c.AssignmentUseCase.Delete(id)

	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, nil)
}

func (c assignmentController) CreateLinkCourseLearningOutcome(ctx *fiber.Ctx) error {
	assignmentId := ctx.Params("assignmentId")
	var payload request.CreateLinkCourseLearningOutcomePayload

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	err := c.AssignmentUseCase.CreateLinkCourseLearningOutcome(assignmentId, payload.CourseLearningOutcomeIds)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusCreated, nil)
}

func (c assignmentController) DeleteLinkCourseLearningOutcome(ctx *fiber.Ctx) error {
	assignmentId := ctx.Params("assignmentId")
	cloId := ctx.Params("cloId")

	err := c.AssignmentUseCase.DeleteLinkCourseLearningOutcome(assignmentId, cloId)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, nil)
}
