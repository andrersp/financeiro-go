package persistence

import (
	"fmt"

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
		fmt.Println(err)
		return
	}
	user = &userData

	return
}

func (u *userRepo) GetUser(userID uint64) (user *entity.User, err error) {

	err = u.db.Debug().First(&user, userID).Error

	return

}

func (u *userRepo) GetUsers() (users []entity.PublicUser, err error) {

	query := u.db.Debug().Model(&entity.User{}).Find(&users)

	if err = query.Error; err != nil {
		return
	}

	return
}

func (u *userRepo) GetUserByEmail(email string) (user *entity.User, err error) {

	err = u.db.Debug().Model(&entity.User{}).Where("email =  ?", email).First(&user).Error
	return
}
