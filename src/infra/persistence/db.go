package persistence

import (
	"github.com/andrersp/financeiro-go/src/domain/entity"
	"github.com/andrersp/financeiro-go/src/domain/repository"
	"gorm.io/gorm"
)

type Repositories struct {
	User repository.UserRepository
	db   *gorm.DB
}

func NewRepositories(db *gorm.DB) (*Repositories, error) {

	return &Repositories{
		db:   db,
		User: NewUserRepository(db),
	}, nil

}

func (r *Repositories) AutoMigrate() error {
	return r.db.AutoMigrate(&entity.User{})

}
