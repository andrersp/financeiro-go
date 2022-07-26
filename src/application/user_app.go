package application

import (
	"github.com/andrersp/financeiro-go/src/domain/entity"
	"github.com/andrersp/financeiro-go/src/domain/repository"
)

type userApp struct {
	repository repository.UserRepository
}

func NewUserApp(repository repository.UserRepository) repository.UserRepository {
	return &userApp{
		repository: repository,
	}
}

func (u *userApp) SaveUser(userData entity.User) (user *entity.User, err error) {
	return
}

func (u *userApp) GetUser(userId uint64) (user *entity.User, err error) {
	return
}

func (u *userApp) GetUsers() (users []entity.User, err error) {
	return
}

func (u *userApp) GetUserByEmail(email string) (user *entity.User, err error) {
	return
}
