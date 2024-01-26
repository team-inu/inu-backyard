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

func (c enrollmentController) GetByID(ctx *fiber.Ctx) error {
	enrollmentID := ctx.Params("enrollmentID")

	enrollment, err := c.EnrollmentUseCase.GetByID(enrollmentID)

	if err != nil {
		return err
	}

	return ctx.JSON(enrollment)
}

func (c enrollmentController) Create(ctx *fiber.Ctx) error {
	var payload request.CreateEnrollmentPayload

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	createdEnrollment, err := c.EnrollmentUseCase.Create(payload.CourseID, payload.StudentID)
	if err != nil {
		return err
	}

	return ctx.JSON(createdEnrollment)
}

func (c enrollmentController) Update(ctx *fiber.Ctx) error {
	enrollmentID := ctx.Params("enrollmentID")

	_, err := c.EnrollmentUseCase.GetByID(enrollmentID)
	if err != nil {
		return err
	}
	var payload request.UpdateEnrollmentPayload

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	err = c.EnrollmentUseCase.Update(enrollmentID, &entity.Enrollment{
		CourseID:  payload.CourseID,
		StudentID: payload.StudentID,
	})
	if err != nil {
		return err
	}

	return ctx.JSON(payload)
}

func (c enrollmentController) Delete(ctx *fiber.Ctx) error {
	enrollmentID := ctx.Params("enrollmentID")

	_, err := c.EnrollmentUseCase.GetByID(enrollmentID)
	if err != nil {
		return err
	}

	err = c.EnrollmentUseCase.Delete(enrollmentID)
	if err != nil {
		return err
	}

	return ctx.SendStatus(200)
}
