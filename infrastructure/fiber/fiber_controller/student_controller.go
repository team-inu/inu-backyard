package fiber_controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team-inu/inu-backyard/entity"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/request"
	"github.com/team-inu/inu-backyard/infrastructure/validator"
)

type studentController struct {
	StudentUsecase entity.StudentUsecase
	Validator      validator.Validator
}

func NewStudentController(studentUsecase entity.StudentUsecase) *studentController {
	return &studentController{
		StudentUsecase: studentUsecase,
		Validator:      validator.NewPayloadValidator(),
	}
}

func (c studentController) GetAll(ctx *fiber.Ctx) error {
	students, err := c.StudentUsecase.GetAll()
	if err != nil {
		return err
	}

	return ctx.JSON(students)
}

func (c studentController) GetByID(ctx *fiber.Ctx) error {
	studentID := ctx.Params("studentID")

	student, _ := c.StudentUsecase.GetByID(studentID)

	return ctx.JSON(student)
}

func (c studentController) Create(ctx *fiber.Ctx) error {
	var student request.CreateStudentRequestBody
	err := ctx.BodyParser(&student)
	if err != nil {
		return err
	}

	validationErrors := c.Validator.Struct(student)
	if len(validationErrors) > 0 {
		return ctx.JSON(validationErrors)
	}

	createdStudent, err := c.StudentUsecase.Create(student.KmuttID, student.Name, student.FirstName, student.LastName)
	if err != nil {
		return err
	}

	return ctx.JSON(createdStudent)
}
