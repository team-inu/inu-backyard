package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team-inu/inu-backyard/entity"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/middleware"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/request"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/response"
	"github.com/team-inu/inu-backyard/internal/validator"
)

type userController struct {
	userUseCase entity.UserUseCase
	authUseCase entity.AuthUseCase
	Validator   validator.PayloadValidator
}

func NewUserController(validator validator.PayloadValidator, userUseCase entity.UserUseCase, authUseCase entity.AuthUseCase) *userController {
	return &userController{
		userUseCase: userUseCase,
		authUseCase: authUseCase,
		Validator:   validator,
	}
}

func (c userController) GetAll(ctx *fiber.Ctx) error {
	users, err := c.userUseCase.GetAll()
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, users)
}

func (c userController) GetById(ctx *fiber.Ctx) error {
	userId := ctx.Params("userId")

	user, err := c.userUseCase.GetById(userId)
	if err != nil {
		return err
	}

	if user == nil {
		return response.NewSuccessResponse(ctx, fiber.StatusNotFound, user)
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, user)
}

func (c userController) Create(ctx *fiber.Ctx) error {
	var payload request.CreateUserPayload
	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	user := middleware.GetUserFromCtx(ctx)
	if !user.IsRoles([]entity.UserRole{entity.UserRoleHeadOfCurriculum}) {
		return response.NewErrorResponse(ctx, fiber.StatusUnauthorized, nil)
	}

	err := c.userUseCase.Create(payload.FirstName, payload.LastName, payload.Email, payload.Password, payload.Role)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusCreated, nil)
}

func (c userController) CreateMany(ctx *fiber.Ctx) error {
	var payload request.CreateBulkUserPayload
	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	newUsers := make([]entity.User, 0, len(payload.Users))

	for _, user := range payload.Users {
		newUsers = append(newUsers, entity.User{
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Role:      user.Role,
			Password:  user.Password,
		})
	}

	err := c.userUseCase.CreateMany(newUsers)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusCreated, nil)
}

func (c userController) Update(ctx *fiber.Ctx) error {
	var payload request.UpdateUserPayload
	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	targetUserId := ctx.Params("userId")

	user := middleware.GetUserFromCtx(ctx)
	if !user.IsRoles([]entity.UserRole{entity.UserRoleHeadOfCurriculum}) && user.Id != targetUserId {
		return response.NewErrorResponse(ctx, fiber.StatusUnauthorized, nil)
	}

	err := c.userUseCase.Update(targetUserId, &entity.User{
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

func (c userController) Delete(ctx *fiber.Ctx) error {
	targetUserId := ctx.Params("userId")

	user := middleware.GetUserFromCtx(ctx)
	if !user.IsRoles([]entity.UserRole{entity.UserRoleHeadOfCurriculum}) && user.Id != targetUserId {
		return response.NewErrorResponse(ctx, fiber.StatusUnauthorized, nil)
	}

	err := c.userUseCase.Delete(targetUserId)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, nil)
}

func (c userController) ChangePassword(ctx *fiber.Ctx) error {
	var payload request.ChangePasswordPayload
	if ok, err := c.Validator.Validate(&payload, ctx); !ok {
		return err
	}

	targetUserId := ctx.Params("userId")

	user := middleware.GetUserFromCtx(ctx)
	if !user.IsRoles([]entity.UserRole{entity.UserRoleHeadOfCurriculum}) && user.Id != targetUserId {
		return response.NewErrorResponse(ctx, fiber.StatusUnauthorized, nil)
	}

	err := c.authUseCase.ChangePassword(targetUserId, payload.OldPassword, payload.NewPassword)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, nil)
}
