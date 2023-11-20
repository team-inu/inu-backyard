package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team-inu/inu-backyard/entity"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/request"
	"github.com/team-inu/inu-backyard/internal/validator"
)

type courseController struct {
	CourseUsecase entity.CourseUsecase
	Validator     validator.PayloadValidator
}

func NewCourseController(courseUsecase entity.CourseUsecase) *courseController {
	return &courseController{
		CourseUsecase: courseUsecase,
		Validator:     validator.NewPayloadValidator(),
	}
}

func (c courseController) GetAll(ctx *fiber.Ctx) error {
	courses, err := c.CourseUsecase.GetAll()
	if err != nil {
		return err
	}

	return ctx.JSON(courses)
}

func (c courseController) GetByID(ctx *fiber.Ctx) error {
	courseID := ctx.Params("courseID")

	course, err := c.CourseUsecase.GetByID(courseID)
	if err != nil {
		return err
	}

	return ctx.JSON(course)
}

func (c courseController) Create(ctx *fiber.Ctx) error {
	var payload request.CreateCourseRequestPayload
	err := ctx.BodyParser(&payload)
	if err != nil {
		return err
	}

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	createdCourse, err := c.CourseUsecase.Create(payload.Name, payload.Code, payload.Year, payload.LecturerID)
	if err != nil {
		return err
	}

	return ctx.JSON(createdCourse)
}

func (c courseController) Delete(ctx *fiber.Ctx) error {
	courseID := ctx.Params("courseID")

	_, err := c.CourseUsecase.GetByID(courseID)
	if err != nil {
		return err
	}

	err = c.CourseUsecase.Delete(courseID)
	if err != nil {
		return err
	}

	return ctx.SendStatus(200)
}
