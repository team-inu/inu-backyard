package main

import fiber_handler "github.com/team-inu/inu-backyard/infrastructure/fiber"

func main() {
	fiberServer := fiber_handler.NewFiberServer()

	fiberServer.Run()
}
