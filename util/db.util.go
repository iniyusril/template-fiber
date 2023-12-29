package util

import (
	"template-fiber/db"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetDbContext(c *fiber.Ctx) (*gorm.DB, error) {
	dbctx, ok := c.Locals("db").(*gorm.DB)
	if !ok {
		dbctx = db.Get()
	}
	dbctx.WithContext(c.Context())
	return dbctx, nil
}
