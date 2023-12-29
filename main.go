package main

import (
	"template-fiber/factory"
	"template-fiber/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	f := factory.NewFactory()
	route := routes.NewRoutes(app, f)
	route.RegisterRoutes()

	app.Listen((":9000"))
}
