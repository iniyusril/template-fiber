package routes

import (
	"template-fiber/controller"
	"template-fiber/factory"

	"github.com/gofiber/fiber/v2"
)

type Routes struct {
	app               *fiber.App
	ExampleController *controller.ExampleController
}

func NewRoutes(app *fiber.App, f *factory.Factory) Routes {
	return Routes{
		app:               app,
		ExampleController: controller.NewExampleController(f),
	}
}

func (r Routes) RegisterRoutes() {
	//register routes here
	app := r.app

	exampleGroup := app.Group("v1/example")
	{
		exampleGroup.Post("", r.ExampleController.Create)
	}

}
