package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team-inu/inu-backyard/entity"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/request"
	"github.com/team-inu/inu-backyard/internal/validator"
)

type studentController struct {
	StudentUseCase entity.StudentUseCase
	Validator      validator.PayloadValidator
}

func NewStudentController(studentUseCase entity.StudentUseCase) *studentController {
	return &studentController{
		StudentUseCase: studentUseCase,
		Validator:      validator.NewPayloadValidator(),
	}
}

func (c studentController) GetAll(ctx *fiber.Ctx) error {
	students, err := c.StudentUseCase.GetAll()
	if err != nil {
		return err
	}

	return ctx.JSON(students)
}

func (c studentController) GetByID(ctx *fiber.Ctx) error {
	studentID := ctx.Params("studentID")

	student, err := c.StudentUseCase.GetByID(studentID)

	if err != nil {
		return err
	}

	return ctx.JSON(student)
}

func (c studentController) GetStudents(ctx *fiber.Ctx) error {
	var payload request.GetStudentsByParamsPayload

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	student, err := c.StudentUseCase.GetByParams(&entity.Student{
		ProgrammeID:    payload.ProgrammeID,
		DepartmentName: payload.DepartmentName,
		Year:           payload.Year,
	}, -1, -1)

	if err != nil {
		return err
	}

	return ctx.JSON(student)
}

func (c studentController) Create(ctx *fiber.Ctx) error {
	var payload request.CreateStudentPayload

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	err := c.StudentUseCase.Create(&entity.Student{
		ID:             payload.KmuttID,
		Name:           payload.Name,
		FirstName:      payload.FirstName,
		LastName:       payload.LastName,
		ProgrammeID:    payload.ProgrammeID,
		DepartmentName: payload.DepartmentName,
		GPAX:           payload.GPAX,
		MathGPA:        payload.MathGPA,
		EngGPA:         payload.EngGPA,
		SciGPA:         payload.SciGPA,
		School:         payload.School,
		Year:           payload.Year,
		Admission:      payload.Admission,
		Remark:         payload.Remark,
	})
	if err != nil {
		return err
	}

	return ctx.JSON(payload)
}

func (c studentController) CreateMany(ctx *fiber.Ctx) error {
	var payload request.CreateBulkStudentsPayload

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	newStudent := []entity.Student{}

	for _, student := range payload.Students {
		newStudent = append(newStudent, entity.Student{
			ID:             student.KmuttID,
			Name:           student.Name,
			FirstName:      student.FirstName,
			LastName:       student.LastName,
			ProgrammeID:    student.ProgrammeID,
			DepartmentName: student.DepartmentName,
			GPAX:           student.GPAX,
			MathGPA:        student.MathGPA,
			EngGPA:         student.EngGPA,
			SciGPA:         student.SciGPA,
			School:         student.School,
			Year:           student.Year,
			Admission:      student.Admission,
			Remark:         student.Remark,
		})
	}

	err := c.StudentUseCase.CreateMany(newStudent)
	if err != nil {
		return err
	}

	return ctx.JSON(payload)
}
