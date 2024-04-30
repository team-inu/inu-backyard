package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team-inu/inu-backyard/entity"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/response"
	"github.com/team-inu/inu-backyard/internal/validator"
)

type coursePortfolioController struct {
	coursePortfolioUseCase entity.CoursePortfolioUseCase
	Validator              validator.PayloadValidator
}

func NewCoursePortfolioController(validator validator.PayloadValidator, coursePortfolioUseCase entity.CoursePortfolioUseCase) *coursePortfolioController {
	return &coursePortfolioController{
		coursePortfolioUseCase: coursePortfolioUseCase,
		Validator:              validator,
	}
}

func (c coursePortfolioController) Generate(ctx *fiber.Ctx) error {
	courseId := ctx.Params("courseId")

	coursePortfolio, err := c.coursePortfolioUseCase.Generate(courseId)
	if err != nil {
		return err
	}
	return response.NewSuccessResponse(ctx, fiber.StatusOK, coursePortfolio)
}

func (c coursePortfolioController) GetCloPassingStudentsByCourseId(ctx *fiber.Ctx) error {
	courseId := ctx.Params("courseId")

	records, err := c.coursePortfolioUseCase.GetCloPassingStudentsByCourseId(courseId)
	if err != nil {
		return err
	}
	return response.NewSuccessResponse(ctx, fiber.StatusOK, records)
}
