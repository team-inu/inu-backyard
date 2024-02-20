package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team-inu/inu-backyard/entity"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/request"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/response"
	"github.com/team-inu/inu-backyard/internal/validator"
)

type studentController struct {
	studentUseCase entity.StudentUseCase
	Validator      validator.PayloadValidator
}

func NewStudentController(validator validator.PayloadValidator, studentUseCase entity.StudentUseCase) *studentController {
	return &studentController{
		studentUseCase: studentUseCase,
		Validator:      validator,
	}
}

func (c studentController) GetAll(ctx *fiber.Ctx) error {
	students, err := c.studentUseCase.GetAll()
	if err != nil {
		return err
	}

	return ctx.JSON(students)
}

func (c studentController) GetById(ctx *fiber.Ctx) error {
	studentId := ctx.Params("studentId")

	student, err := c.studentUseCase.GetById(studentId)

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

	student, err := c.studentUseCase.GetByParams(&entity.Student{
		ProgrammeId:    payload.ProgrammeId,
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

	err := c.studentUseCase.Create(&entity.Student{
		Id:             payload.KmuttId,
		Name:           payload.Name,
		FirstName:      payload.FirstName,
		LastName:       payload.LastName,
		ProgrammeId:    payload.ProgrammeId,
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
			Id:             student.KmuttId,
			Name:           student.Name,
			FirstName:      student.FirstName,
			LastName:       student.LastName,
			ProgrammeId:    student.ProgrammeId,
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

	err := c.studentUseCase.CreateMany(newStudent)
	if err != nil {
		return err
	}

	return ctx.JSON(payload)
}

func (c studentController) Update(ctx *fiber.Ctx) error {
	var payload request.UpdateStudentPayload

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	id := ctx.Params("studentId")

	err := c.studentUseCase.Update(id, &entity.Student{
		Id:             payload.KmuttId,
		Name:           payload.Name,
		FirstName:      payload.FirstName,
		LastName:       payload.LastName,
		ProgrammeId:    payload.ProgrammeId,
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

	return response.NewSuccessResponse(ctx, fiber.StatusOK, nil)
}

func (c studentController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("studentId")

	err := c.studentUseCase.Delete(id)

	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, nil)
}
