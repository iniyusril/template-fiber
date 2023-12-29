package repositories

import (
	"template-fiber/db"
	"template-fiber/entities"

	"github.com/gofiber/fiber/v2"
)

type ExampleRepository struct {
}

func NewExampleRepository() *ExampleRepository {
	return &ExampleRepository{}
}

func (r *ExampleRepository) Create(c *fiber.Ctx, data *entities.Example) error {
	db := db.GetDbContext(c)
	return db.Create(data).Error
}
