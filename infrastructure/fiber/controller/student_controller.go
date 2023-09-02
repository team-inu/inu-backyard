package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/oklog/ulid/v2"
	"github.com/team-inu/inu-backyard/entity"
)

type studentController struct {
	StudentUsecase entity.StudentUsecase
}

func NewStudentController(studentUsecase entity.StudentUsecase) *studentController {
	return &studentController{StudentUsecase: studentUsecase}
}

func (s studentController) GetAll(ctx *fiber.Ctx) error {
	students, err := s.StudentUsecase.GetAll()
	if err != nil {
		return err
	}

	return ctx.JSON(students)
}

func (s studentController) GetByID(ctx *fiber.Ctx) error {
	studentID := ctx.Params("studentID")

	studentUUID, _ := ulid.Parse(studentID)
	student, _ := s.StudentUsecase.GetByID(studentUUID)

	return ctx.JSON(student)
}

// read body and create student
func (s studentController) Create(ctx *fiber.Ctx) error {
	student := new(entity.Student)

	if err := ctx.BodyParser(student); err != nil {
		fmt.Println(err)
		// return err
	}

	if err := s.StudentUsecase.Create(student); err != nil {
		return err
	}

	return ctx.JSON(student)
}
