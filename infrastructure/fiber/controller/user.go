package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team-inu/inu-backyard/entity"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/request"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/response"
	"github.com/team-inu/inu-backyard/internal/validator"
)

type userController struct {
	userUseCase entity.UserUseCase
	Validator   validator.PayloadValidator
}

func NewUserController(validator validator.PayloadValidator, userUseCase entity.UserUseCase) *userController {
	return &userController{
		userUseCase: userUseCase,
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

	err := c.userUseCase.Create(payload.FirstName, payload.LastName, payload.Email, payload.Password)

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

	id := ctx.Params("userId")

	err := c.userUseCase.Update(id, &entity.User{
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
	id := ctx.Params("userId")

	err := c.userUseCase.Delete(id)

	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ctx, fiber.StatusOK, nil)
}
