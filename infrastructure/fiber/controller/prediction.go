package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team-inu/inu-backyard/entity"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/response"
	"github.com/team-inu/inu-backyard/internal/validator"
)

type predictionController struct {
	predictionUseCase entity.PredictionUseCase
	Validator         validator.PayloadValidator
}

func NewPredictionController(validator validator.PayloadValidator, predictionUseCase entity.PredictionUseCase) *predictionController {
	return &predictionController{
		predictionUseCase: predictionUseCase,
		Validator:         validator,
	}
}

func (c predictionController) Train(ctx *fiber.Ctx) error {
	id, err := c.predictionUseCase.CreatePrediction()
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, id)
}
