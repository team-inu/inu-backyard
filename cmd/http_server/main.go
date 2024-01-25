package main

import (
	"github.com/team-inu/inu-backyard/infrastructure/fiber"
	"github.com/team-inu/inu-backyard/internal/config"
)

func main() {
	var fiberConfig config.FiberServerConfig

	config.SetConfig(&fiberConfig)
	config.PrintConfig()

	fiberServer := fiber.NewFiberServer()

	fiberServer.Run(fiberConfig)
}
