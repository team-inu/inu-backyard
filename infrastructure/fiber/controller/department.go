package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team-inu/inu-backyard/entity"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/request"
	"github.com/team-inu/inu-backyard/internal/validator"
)

type DepartmentController struct {
	departmentUseCase entity.DepartmentUseCase
	validator         validator.PayloadValidator
}

func NewDepartmentController(departmentUseCase entity.DepartmentUseCase) *DepartmentController {
	return &DepartmentController{
		departmentUseCase: departmentUseCase,
		validator:         validator.NewPayloadValidator(),
	}
}

func (c DepartmentController) Create(ctx *fiber.Ctx) error {
	var payload request.CreateDepartmentRequestBody

	if ok, err := c.validator.Validate(&payload, ctx); !ok {
		return err
	}

	err := c.departmentUseCase.Create(&entity.Department{
		Name:        payload.Name,
		FacultyName: payload.FacultyName,
	})

	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"message": "success created",
		"status":  200,
	})
}

func (c DepartmentController) Delete(ctx *fiber.Ctx) error {
	var payload request.DeleteDepartmentRequestBody

	if ok, err := c.validator.Validate(&payload, ctx); !ok {
		return err
	}

	err := c.departmentUseCase.Delete(payload.Name)
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"message": "success deleted",
		"status":  200,
	})
}

func (c DepartmentController) GetAll(ctx *fiber.Ctx) error {
	departments, err := c.departmentUseCase.GetAll()
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"status":  200,
		"data":    departments,
	})
}

func (c DepartmentController) GetByName(ctx *fiber.Ctx) error {
	departmentID := ctx.Params("departmentID")

	department, err := c.departmentUseCase.GetByName(departmentID)

	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"status":  200,
		"data":    department,
	})
}

func (c DepartmentController) Update(ctx *fiber.Ctx) error {
	var payload request.UpdateDepartmentRequestBody

	if ok, err := c.validator.Validate(&payload, ctx); !ok {
		return err
	}

	err := c.departmentUseCase.Update(&entity.Department{
		Name:        payload.Name,
		FacultyName: payload.FacultyName,
	}, payload.NewName)

	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"message": "success updated",
		"status":  200,
	})
}
