package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/request"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/response"
)

func (c assignmentController) GetGroupByCourseId(ctx *fiber.Ctx) error {
	courseId := ctx.Params("courseId")

	assignmentGroups, err := c.AssignmentUseCase.GetGroupByCourseId(courseId)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, assignmentGroups)
}

func (c assignmentController) CreateGroup(ctx *fiber.Ctx) error {
	var payload request.CreateAssignmentGroupPayload
	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	err := c.AssignmentUseCase.CreateGroup(payload.Name, payload.CourseId)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusCreated, nil)
}

func (c assignmentController) UpdateGroup(ctx *fiber.Ctx) error {
	var payload request.UpdateAssignmentGroupPayload
	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	id := ctx.Params("assignmentGroupId")

	err := c.AssignmentUseCase.UpdateGroup(id, payload.Name)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, nil)
}

func (c assignmentController) DeleteGroup(ctx *fiber.Ctx) error {
	id := ctx.Params("assignmentGroupId")

	err := c.AssignmentUseCase.DeleteGroup(id)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, nil)
}
