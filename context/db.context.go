package context

import (
	"context"

	"gorm.io/gorm"
)

type DbContext struct {
	DB  *gorm.DB
	Ctx context.Context
}
