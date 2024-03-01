package entity

import "github.com/gofiber/fiber/v2"

type AuthUseCase interface {
	Authenticate(header string) (*User, error)
	SignIn(email string, password string, ipAddress string, userAgent string) (*fiber.Cookie, error)
	SignOut(header string) (*fiber.Cookie, error)
}
