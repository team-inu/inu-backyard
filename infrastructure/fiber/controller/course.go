package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team-inu/inu-backyard/entity"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/middleware"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/request"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/response"
	"github.com/team-inu/inu-backyard/internal/validator"
)

type courseController struct {
	courseUseCase entity.CourseUseCase
	Validator     validator.PayloadValidator
}

func NewCourseController(validator validator.PayloadValidator, courseUseCase entity.CourseUseCase) *courseController {
	return &courseController{
		courseUseCase: courseUseCase,
		Validator:     validator,
	}
}

func (c courseController) GetAll(ctx *fiber.Ctx) error {
	user := middleware.GetUserFromCtx(ctx)

	var courses []entity.Course
	var err error

	if user.IsRoles([]entity.UserRole{entity.UserRoleHeadOfCurriculum, entity.UserRoleModerator, entity.UserRoleTABEEManager}) {
		courses, err = c.courseUseCase.GetAll()
	} else {
		courses, err = c.courseUseCase.GetByUserId(user.Id)
	}

	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, courses)
}

func (c courseController) GetById(ctx *fiber.Ctx) error {
	courseId := ctx.Params("courseId")

	course, err := c.courseUseCase.GetById(courseId)
	if err != nil {
		return err
	}

	if course == nil {
		return response.NewSuccessResponse(ctx, fiber.StatusNotFound, course)
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, course)
}

func (c courseController) Create(ctx *fiber.Ctx) error {
	var payload request.CreateCourseRequestPayload

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	user := middleware.GetUserFromCtx(ctx)

	err := c.courseUseCase.Create(
		*user,
		payload.SemesterId,
		payload.UserId,
		payload.Name,
		payload.Code,
		payload.Curriculum,
		payload.Description,
		payload.ExpectedPassingCloPercentage,
		payload.CriteriaGrade,
	)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusCreated, nil)
}

func (c courseController) Update(ctx *fiber.Ctx) error {
	var payload request.UpdateCourseRequestPayload

	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	id := ctx.Params("courseId")

	user := middleware.GetUserFromCtx(ctx)

	err := c.courseUseCase.Update(
		*user,
		id,
		payload.Name,
		payload.Code,
		payload.Curriculum,
		payload.Description,
		payload.ExpectedPassingCloPercentage,
		payload.CriteriaGrade,
		*payload.IsPortfolioCompleted,
	)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, nil)
}

func (c courseController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("courseId")

	user := middleware.GetUserFromCtx(ctx)

	err := c.courseUseCase.Delete(*user, id)

	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, nil)
}
