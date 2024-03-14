package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team-inu/inu-backyard/entity"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/request"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/response"
	"github.com/team-inu/inu-backyard/internal/validator"
)

type scoreController struct {
	ScoreUseCase entity.ScoreUseCase
	Validator    validator.PayloadValidator
}

func NewScoreController(validator validator.PayloadValidator, scoreUseCase entity.ScoreUseCase) *scoreController {
	return &scoreController{
		ScoreUseCase: scoreUseCase,
		Validator:    validator,
	}
}

func (c scoreController) GetAll(ctx *fiber.Ctx) error {
	scores, err := c.ScoreUseCase.GetAll()
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, scores)
}

func (c scoreController) GetById(ctx *fiber.Ctx) error {
	scoreId := ctx.Params("scoreId")

	score, err := c.ScoreUseCase.GetById(scoreId)
	if err != nil {
		return err
	}

	if score == nil {
		return response.NewSuccessResponse(ctx, fiber.StatusNotFound, score)
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, score)
}

func (c scoreController) GetByAssignmentId(ctx *fiber.Ctx) error {
	assignmentId := ctx.Params("assignmentId")

	assignmentScore, err := c.ScoreUseCase.GetByAssignmentId(assignmentId)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, assignmentScore)
}

func (c scoreController) CreateMany(ctx *fiber.Ctx) error {
	var payload request.BulkCreateScoreRequestPayload
	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	user := ctx.Locals("user").(*entity.User)

	err := c.ScoreUseCase.CreateMany(
		user.Id,
		payload.AssignmentId,
		payload.StudentScores,
	)

	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusCreated, nil)
}

func (c scoreController) Delete(ctx *fiber.Ctx) error {
	scoreId := ctx.Params("scoreId")

	_, err := c.ScoreUseCase.GetById(scoreId)
	if err != nil {
		return err
	}

	err = c.ScoreUseCase.Delete(scoreId)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, nil)
}

func (c scoreController) Update(ctx *fiber.Ctx) error {
	scoreId := ctx.Params("scoreId")

	_, err := c.ScoreUseCase.GetById(scoreId)
	if err != nil {
		return err
	}
	var payload request.UpdateScoreRequestPayload

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	err = c.ScoreUseCase.Update(scoreId, payload.Score)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, nil)
}
