package controller

import (
	"net/http"
	"template-fiber/dto"
	"template-fiber/dto/response"
	"template-fiber/factory"
	"template-fiber/services"

	"github.com/gofiber/fiber/v2"
)

type ExampleController struct {
	ExampleService *services.ExampleService
}

func NewExampleController(f *factory.Factory) *ExampleController {
	return &ExampleController{
		ExampleService: services.NewExampleService(f),
	}
}

func (s *ExampleController) Create(c *fiber.Ctx) error {

	var payload = new(dto.ExampleRequest)

	if err := c.BodyParser(payload); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.ErrorBuilder(err))
	}

	data, err := s.ExampleService.Create(c, payload)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.ErrorBuilder(err))
	}

	return c.JSON(response.SuccessBuilder(data))
}
