package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team-inu/inu-backyard/entity"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/request"
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

func (c predictionController) Predict(ctx *fiber.Ctx) error {
	var payload request.PredictPayload
	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	prediction, err := c.predictionUseCase.CreatePrediction(entity.PredictionRequirements{
		ProgrammeName: payload.ProgrammeName,
		OldGPAX:       payload.GPAX,
		MathGPA:       payload.MathGPA,
		EngGPA:        payload.EngGPA,
		SciGPA:        payload.SciGPA,
		School:        payload.School,
		Admission:     payload.Admission,
	})
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, prediction)
}
