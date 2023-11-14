package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team-inu/inu-backyard/entity"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/request"
	"github.com/team-inu/inu-backyard/internal/validator"
)

type studentController struct {
	StudentUseCase entity.StudentUseCase
	Validator      validator.Validator
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

	student, _ := c.StudentUseCase.GetByID(studentID)

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

	err = c.StudentUseCase.Create(&entity.Student{
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
	if err != nil {
		return err
	}

	return ctx.JSON(student)
}

func (c studentController) CreateMany(ctx *fiber.Ctx) error {
	var body request.CreateBulkStudentsRequestBody
	err := ctx.BodyParser(&body)
	if err != nil {
		return err
	}

	validationErrors := c.Validator.Struct(body)
	if len(validationErrors) > 0 {
		return ctx.JSON(validationErrors)
	}

	newStudent := []entity.Student{}

	for _, student := range body.Students {
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

	err = c.StudentUseCase.CreateMany(newStudent)
	if err != nil {
		return err
	}

	return ctx.JSON(body)
}
