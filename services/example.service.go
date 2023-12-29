package services

import (
	"template-fiber/db"
	"template-fiber/dto"
	"template-fiber/entities"
	"template-fiber/factory"
	"template-fiber/repositories"

	"github.com/gofiber/fiber/v2"
	"github.com/stroiman/go-automapper"
	"gorm.io/gorm"
)

type ExampleService struct {
	ExampleRepository repositories.ExampleRepository
}

func NewExampleService(f *factory.Factory) *ExampleService {
	return &ExampleService{
		ExampleRepository: f.ExampleRepository,
	}
}

func (s *ExampleService) Create(c *fiber.Ctx, payload *dto.ExampleRequest) (*entities.Example, error) {
	var entity = new(entities.Example)

	automapper.MapLoose(payload, &entity)

	err := db.Get().Transaction(func(tx *gorm.DB) error {
		c.Locals("db", tx)

		err := s.ExampleRepository.Create(c, entity)
		if err != nil {
			return err
		}

		// return errors.New("throw error")
		return nil
	})

	return entity, err
}
