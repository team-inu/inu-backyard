package main

import (
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

	fiberServer := fiber.NewFiberServer(
		fiberConfig,
		gormDB,
		zapLogger,
	)

	fiberServer.Run()
}
