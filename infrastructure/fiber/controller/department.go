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

func NewDepartmentController(validator validator.PayloadValidator, departmentUseCase entity.DepartmentUseCase) *DepartmentController {
	return &DepartmentController{
		departmentUseCase: departmentUseCase,
		validator:         validator,
	}
}

func (c DepartmentController) Create(ctx *fiber.Ctx) error {
	var payload request.CreateDepartmentRequestPayload

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

	return ctx.JSON(payload)
}

func (c DepartmentController) Delete(ctx *fiber.Ctx) error {
	departmentName := ctx.Params("departmentName")

	_, err := c.departmentUseCase.GetByName(departmentName)

	if err != nil {
		return err
	}

	err = c.departmentUseCase.Delete(departmentName)

	if err != nil {
		return err
	}

	return ctx.SendStatus(200)
}

func (c DepartmentController) GetAll(ctx *fiber.Ctx) error {
	departments, err := c.departmentUseCase.GetAll()
	if err != nil {
		return err
	}

	return ctx.JSON(departments)
}

func (c DepartmentController) GetByName(ctx *fiber.Ctx) error {
	departmentName := ctx.Params("departmentName")

	department, err := c.departmentUseCase.GetByName(departmentName)

	if err != nil {
		return err
	}

	return ctx.JSON(department)
}

func (c DepartmentController) Update(ctx *fiber.Ctx) error {
	departmentName := ctx.Params("departmentName")

	_, err := c.departmentUseCase.GetByName(departmentName)

	if err != nil {
		return err
	}

	var payload request.UpdateDepartmentRequestPayload

	if ok, err := c.validator.Validate(&payload, ctx); !ok {
		return err
	}

	err = c.departmentUseCase.Update(&entity.Department{
		Name:        departmentName,
		FacultyName: payload.FacultyName,
	}, payload.NewName)

	if err != nil {
		return err
	}

	return ctx.SendStatus(200)
}
