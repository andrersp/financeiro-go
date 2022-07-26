package persistence

import (
	"github.com/andrersp/financeiro-go/src/domain/entity"
	"github.com/andrersp/financeiro-go/src/domain/repository"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) SaveUser(userData entity.User) (user *entity.User, err error) {

	err = u.db.Debug().Create(&userData).Error
	if err != nil {
		return
	}
	user = &userData

	return
}

func (u *userRepo) GetUser(userID uint64) (user *entity.User, err error) {
	return
}

func (u *userRepo) GetUsers() (uers []entity.User, err error) {
	return
}

func (u *userRepo) GetUserByEmail(email string) (user *entity.User, err error) {
	return
}
