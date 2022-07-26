package persistence

import (
	"gorm.io/gorm"
)

type Repositories struct {
	db *gorm.DB
}
