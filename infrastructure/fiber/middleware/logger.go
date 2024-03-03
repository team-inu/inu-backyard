package middleware

import (
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
)

func NewLogger(config fiberzap.Config) fiber.Handler {
	return fiberzap.New(config)
}
