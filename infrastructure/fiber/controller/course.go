package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team-inu/inu-backyard/entity"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/request"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/response"
	"github.com/team-inu/inu-backyard/internal/validator"
)

type courseController struct {
	courseUsecase entity.CourseUsecase
	Validator     validator.PayloadValidator
}

func NewCourseController(validator validator.PayloadValidator, courseUsecase entity.CourseUsecase) *courseController {
	return &courseController{
		courseUsecase: courseUsecase,
		Validator:     validator,
	}
}

func (c courseController) GetAll(ctx *fiber.Ctx) error {
	courses, err := c.courseUsecase.GetAll()
	if err != nil {
		return err
	}

	return ctx.JSON(courses)
}

func (c courseController) GetByID(ctx *fiber.Ctx) error {
	courseID := ctx.Params("courseID")

	course, err := c.courseUsecase.GetByID(courseID)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusCreated, course)
}

func (c courseController) Create(ctx *fiber.Ctx) error {
	var payload request.CreateCourseRequestPayload

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	err := c.courseUsecase.Create(payload.Name, payload.Code, payload.SemesterID, payload.LecturerID)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusCreated, nil)
}

func (c courseController) Update(ctx *fiber.Ctx) error {
	var payload request.UpdateCourseRequestPayload

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	id := ctx.Params("courseID")

	err := c.courseUsecase.Update(id, &entity.Course{
		Name:       payload.Name,
		Code:       payload.Code,
		SemesterID: payload.SemesterID,
		LecturerID: payload.LecturerID,
	})

	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, nil)
}

func (c courseController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("courseID")

	err := c.courseUsecase.Delete(id)

	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, nil)
}
