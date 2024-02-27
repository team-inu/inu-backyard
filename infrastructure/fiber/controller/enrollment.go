package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team-inu/inu-backyard/entity"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/request"
	"github.com/team-inu/inu-backyard/internal/validator"
)

type enrollmentController struct {
	EnrollmentUseCase entity.EnrollmentUseCase
	Validator         validator.PayloadValidator
}

func NewEnrollmentController(validator validator.PayloadValidator, enrollmentUseCase entity.EnrollmentUseCase) *enrollmentController {
	return &enrollmentController{
		EnrollmentUseCase: enrollmentUseCase,
		Validator:         validator,
	}
}

func (c enrollmentController) GetAll(ctx *fiber.Ctx) error {
	enrollments, err := c.EnrollmentUseCase.GetAll()
	if err != nil {
		return err
	}

	return ctx.JSON(enrollments)
}

func (c enrollmentController) GetById(ctx *fiber.Ctx) error {
	enrollmentId := ctx.Params("enrollmentId")

	enrollment, err := c.EnrollmentUseCase.GetById(enrollmentId)

	if err != nil {
		return err
	}

	return ctx.JSON(enrollment)
}

func (c enrollmentController) GetByCourseId(ctx *fiber.Ctx) error {
	enrollmentId := ctx.Params("courseId")

	enrollments, err := c.EnrollmentUseCase.GetByCourseId(enrollmentId)

	if err != nil {
		return err
	}

	return ctx.JSON(enrollments)
}

func (c enrollmentController) Create(ctx *fiber.Ctx) error {
	var payload request.CreateEnrollmentsPayload

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	err := c.EnrollmentUseCase.CreateMany(payload.CourseId, payload.Status, payload.StudentIds)
	if err != nil {
		return err
	}

	return ctx.SendStatus(201)
}

func (c enrollmentController) Update(ctx *fiber.Ctx) error {
	enrollmentId := ctx.Params("enrollmentId")

	var payload request.UpdateEnrollmentPayload

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	err := c.EnrollmentUseCase.Update(enrollmentId, payload.Status)
	if err != nil {
		return err
	}

	return ctx.JSON(payload)
}

func (c enrollmentController) Delete(ctx *fiber.Ctx) error {
	enrollmentId := ctx.Params("enrollmentId")

	_, err := c.EnrollmentUseCase.GetById(enrollmentId)
	if err != nil {
		return err
	}

	err = c.EnrollmentUseCase.Delete(enrollmentId)
	if err != nil {
		return err
	}

	return ctx.SendStatus(200)
}
