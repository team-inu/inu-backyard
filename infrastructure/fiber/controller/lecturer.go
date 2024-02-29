package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team-inu/inu-backyard/entity"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/request"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/response"
	"github.com/team-inu/inu-backyard/internal/validator"
)

type lecturerController struct {
	lecturerUseCase entity.LecturerUseCase
	Validator       validator.PayloadValidator
}

func NewLecturerController(validator validator.PayloadValidator, lecturerUseCase entity.LecturerUseCase) *lecturerController {
	return &lecturerController{
		lecturerUseCase: lecturerUseCase,
		Validator:       validator,
	}
}

func (c lecturerController) GetAll(ctx *fiber.Ctx) error {
	lecturers, err := c.lecturerUseCase.GetAll()
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, lecturers)
}

func (c lecturerController) GetById(ctx *fiber.Ctx) error {
	lecturerId := ctx.Params("lecturerId")

	lecturer, err := c.lecturerUseCase.GetById(lecturerId)

	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusCreated, lecturer)
}

func (c lecturerController) Create(ctx *fiber.Ctx) error {
	var payload request.CreateLecturerPayload

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	err := c.lecturerUseCase.Create(payload.FirstName, payload.LastName, payload.Email, payload.Password)

	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusCreated, nil)
}

func (c lecturerController) CreateMany(ctx *fiber.Ctx) error {
	var payload request.CreateBulkLecturerPayload

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	newLectuers := make([]entity.Lecturer, 0, len(payload.Lecturers))

	for _, lecturer := range payload.Lecturers {
		newLectuers = append(newLectuers, entity.Lecturer{
			FirstName: lecturer.FirstName,
			LastName:  lecturer.LastName,
			Email:     lecturer.Email,
			Role:      lecturer.Role,
			Password:  lecturer.Password,
		})
	}

	err := c.lecturerUseCase.CreateMany(newLectuers)

	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusCreated, nil)
}

func (c lecturerController) Update(ctx *fiber.Ctx) error {
	var payload request.UpdateLecturerPayload

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	id := ctx.Params("lecturerId")

	err := c.lecturerUseCase.Update(id, &entity.Lecturer{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Role:      payload.Role,
	})

	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, nil)
}

func (c lecturerController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("lecturerId")

	err := c.lecturerUseCase.Delete(id)

	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, nil)
}
