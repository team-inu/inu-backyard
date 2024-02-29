package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/team-inu/inu-backyard/entity"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/request"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/response"
	"github.com/team-inu/inu-backyard/internal/validator"
)

type courseController struct {
	courseUseCase entity.CourseUseCase
	Validator     validator.PayloadValidator
}

func NewCourseController(validator validator.PayloadValidator, courseUseCase entity.CourseUseCase) *courseController {
	return &courseController{
		courseUseCase: courseUseCase,
		Validator:     validator,
	}
}

func (c courseController) GetAll(ctx *fiber.Ctx) error {
	courses, err := c.courseUseCase.GetAll()
	if err != nil {
		return err
	}

	if len(courses) == 0 {
		return response.NewSuccessResponse(ctx, fiber.StatusNotFound, courses)
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, courses)
}

func (c courseController) GetById(ctx *fiber.Ctx) error {
	courseId := ctx.Params("courseId")

	course, err := c.courseUseCase.GetById(courseId)
	if err != nil {
		return err
	}

	if course == nil {
		return response.NewSuccessResponse(ctx, fiber.StatusNotFound, course)
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, course)
}

func (c courseController) Create(ctx *fiber.Ctx) error {
	var payload request.CreateCourseRequestPayload

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	fmt.Println(payload.Description[0])
	err := c.courseUseCase.Create(
		payload.SemesterId,
		payload.LecturerId,
		payload.Name,
		payload.Code,
		payload.Curriculum,
		payload.Description,
		*payload.CriteriaGrade,
	)
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

	id := ctx.Params("courseId")

	err := c.courseUseCase.Update(id, &entity.Course{
		Name:       payload.Name,
		Code:       payload.Code,
		SemesterId: payload.SemesterId,
		LecturerId: payload.LecturerId,
	})

	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, nil)
}

func (c courseController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("courseId")

	err := c.courseUseCase.Delete(id)

	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, nil)
}
