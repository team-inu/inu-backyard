package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team-inu/inu-backyard/entity"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/request"
	"github.com/team-inu/inu-backyard/internal/validator"
)

type scoreController struct {
	ScoreUsecase entity.ScoreUsecase
	Validator    validator.PayloadValidator
}

func NewScoreController(scoreUsecase entity.ScoreUsecase) *scoreController {
	return &scoreController{
		ScoreUsecase: scoreUsecase,
		Validator:    validator.NewPayloadValidator(),
	}
}

func (c scoreController) GetAll(ctx *fiber.Ctx) error {
	scores, err := c.ScoreUsecase.GetAll()
	if err != nil {
		return err
	}

	return ctx.JSON(scores)
}

func (c scoreController) GetByID(ctx *fiber.Ctx) error {
	scoreID := ctx.Params("scoreID")

	score, err := c.ScoreUsecase.GetByID(scoreID)
	if err != nil {
		return err
	}

	return ctx.JSON(score)
}

func (c scoreController) Create(ctx *fiber.Ctx) error {
	var payload request.CreateScoreRequestPayload

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	createdScore, err := c.ScoreUsecase.Create(payload.Score, payload.StudentID, payload.AssessmentID, payload.LecturerID)
	if err != nil {
		return err
	}

	return ctx.JSON(createdScore)
}

func (c scoreController) Delete(ctx *fiber.Ctx) error {
	scoreID := ctx.Params("scoreID")

	_, err := c.ScoreUsecase.GetByID(scoreID)
	if err != nil {
		return err
	}

	err = c.ScoreUsecase.Delete(scoreID)
	if err != nil {
		return err
	}

	return ctx.SendStatus(200)
}

func (c scoreController) Update(ctx *fiber.Ctx) error {
	scoreID := ctx.Params("scoreID")

	_, err := c.ScoreUsecase.GetByID(scoreID)
	if err != nil {
		return err
	}
	var payload request.UpdateScoreRequestPayload

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	err = c.ScoreUsecase.Update(scoreID, payload.Score)
	if err != nil {
		return err
	}

	return ctx.JSON(payload)
}
