package db

import (
	"sync"
	"template-fiber/entities"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var once sync.Once
var err error

func Init() {
	Get()
}

func Get() *gorm.DB {

	once.Do(func() {
		dsn := "host=localhost user=root password=root dbname=test_db port=5432 sslmode=disable TimeZone=Asia/Jakarta"
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		db.AutoMigrate(&entities.Example{})
	})

	return db
}

func GetDbContext(c *fiber.Ctx) *gorm.DB {
	dbctx, ok := c.Locals("db").(*gorm.DB)
	if !ok {
		dbctx = Get()
	}
	dbctx.WithContext(c.Context())
	return dbctx
}
