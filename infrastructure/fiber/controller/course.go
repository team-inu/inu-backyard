package controller

import (
	"fmt"

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
	var course request.CreateCourseRequestBody
	err := ctx.BodyParser(&course)
	if err != nil {
		return err
	}

	err, validationErrors := c.Validator.Validate(course, ctx)
	if err != nil {
		return ctx.Status(400).JSON(validationErrors)
	}

	fmt.Println(course)
	createdCourse, err := c.CourseUsecase.Create(course.Name, course.Code, course.Year, course.LecturerID)
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
