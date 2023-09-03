package main

import "github.com/team-inu/inu-backyard/infrastructure/fiber"

func main() {
	fiberServer := fiber.NewFiberServer()

	fiberServer.Run()
}
