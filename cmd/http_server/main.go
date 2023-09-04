package main

import (
	"github.com/team-inu/inu-backyard/infrastructure/config"
	"github.com/team-inu/inu-backyard/infrastructure/fiber"
)

func main() {
	var fiberConfig fiber.FiberServerConfig

	config.SetConfig(&fiberConfig)
	config.PrintConfig()

	fiberServer := fiber.NewFiberServer()

	fiberServer.Run(fiberConfig)
}
