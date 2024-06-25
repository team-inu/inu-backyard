package main

import (
	"github.com/team-inu/inu-backyard/infrastructure/captcha"
	"github.com/team-inu/inu-backyard/infrastructure/database"
	"github.com/team-inu/inu-backyard/infrastructure/fiber"
	"github.com/team-inu/inu-backyard/internal/config"
	"github.com/team-inu/inu-backyard/internal/logger"
)

func main() {
	var fiberConfig config.FiberServerConfig

	config.SetConfig(&fiberConfig)
	config.PrintConfig()

	zapLogger := logger.NewZapLogger()

	gormDB, err := database.NewGorm(&fiberConfig.Database)
	if err != nil {
		panic(err)
	}

	turnstile := captcha.NewTurnstile(fiberConfig.Client.Auth.Turnstile.SecretKey)

	fiberServer := fiber.NewFiberServer(
		fiberConfig,
		gormDB,
		turnstile,
		zapLogger,
	)

	fiberServer.Run()
}
