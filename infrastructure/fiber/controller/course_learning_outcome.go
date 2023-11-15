package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team-inu/inu-backyard/entity"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/request"
	"github.com/team-inu/inu-backyard/internal/validator"
)

type courseLearningOutcomeController struct {
	CourseLearningOutcomeUsecase entity.CourseLearningOutcomeUsecase
	Validator                    validator.PayloadValidator
}

func NewCourseLearningOutcomeController(courseLearningOutcomeUsecase entity.CourseLearningOutcomeUsecase) *courseLearningOutcomeController {
	return &courseLearningOutcomeController{
		CourseLearningOutcomeUsecase: courseLearningOutcomeUsecase,
		Validator:                    validator.NewPayloadValidator(),
	}
}

func (c courseLearningOutcomeController) GetAll(ctx *fiber.Ctx) error {
	clos, err := c.CourseLearningOutcomeUsecase.GetAll()
	if err != nil {
		return err
	}

	return ctx.JSON(clos)
}

func (c courseLearningOutcomeController) GetByID(ctx *fiber.Ctx) error {
	cloId := ctx.Params("cloId")

	clos, err := c.CourseLearningOutcomeUsecase.GetByID(cloId)
	if err != nil {
		return err
	}

	return ctx.JSON(clos)
}

func (c courseLearningOutcomeController) GetByCourseID(ctx *fiber.Ctx) error {
	courseId := ctx.Params("courseId")

	clos, err := c.CourseLearningOutcomeUsecase.GetByCourseID(courseId)
	if err != nil {
		return err
	}

	return ctx.JSON(clos)
}

func (c courseLearningOutcomeController) Create(ctx *fiber.Ctx) error {
	var clo request.CreateCourseLearningOutcomeBody
	err := ctx.BodyParser(&clo)
	if err != nil {
		return err
	}

	err, validationErrors := c.Validator.Validate(clo, ctx)
	if err != nil {
		return ctx.Status(400).JSON(validationErrors)
	}

	createdClo, err := c.CourseLearningOutcomeUsecase.Create(clo.Code, clo.Description, clo.Weight, clo.SubProgramLearningOutcomeID, clo.ProgramOutcomeID, clo.CourseId, clo.Status)
	if err != nil {
		return err
	}

	return ctx.JSON(createdClo)
}

func (c courseLearningOutcomeController) Delete(ctx *fiber.Ctx) error {
	cloId := ctx.Params("cloId")

	_, err := c.CourseLearningOutcomeUsecase.GetByID(cloId)
	if err != nil {
		return err
	}

	err = c.CourseLearningOutcomeUsecase.Delete(cloId)
	if err != nil {
		return err
	}

	return ctx.SendStatus(200)
}
