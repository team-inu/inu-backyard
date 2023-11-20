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
	var department request.CreateDepartmentRequestBody

	err := ctx.BodyParser(&department)
	if err != nil {
		return err
	}

	err, validationErrors := c.validator.Validate(&department, ctx)
	if err != nil {
		return ctx.Status(400).JSON(validationErrors)
	}

	err = c.departmentUseCase.Create(&entity.Department{
		Name:        department.Name,
		FacultyName: department.FacultyName,
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
	var departmentName request.DeleteDepartmentRequestBody

	err := ctx.BodyParser(&departmentName)
	if err != nil {
		return err
	}

	err, validationErrors := c.validator.Validate(&departmentName, ctx)

	if err != nil {
		return ctx.Status(400).JSON(validationErrors)
	}

	err = c.departmentUseCase.Delete(departmentName.Name)

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

	department, err := c.departmentUseCase.GetByID(departmentID)

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
	var department request.UpdateDepartmentRequestBody

	err := ctx.BodyParser(&department)
	if err != nil {
		return err
	}

	err, validationErrors := c.validator.Validate(&department, ctx)
	if err != nil {
		return ctx.Status(400).JSON(validationErrors)
	}

	err = c.departmentUseCase.Update(&entity.Department{
		Name:        department.Name,
		FacultyName: department.FacultyName,
	}, department.NewName)

	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"message": "success updated",
		"status":  200,
	})
}
