package application

import (
	"github.com/andrersp/financeiro-go/src/domain/entity"
	"github.com/andrersp/financeiro-go/src/domain/repository"
)

type UserApp struct {
	repository repository.UserRepository
}

func NewUserApp(repository repository.UserRepository) repository.UserRepository {
	return &UserApp{
		repository: repository,
	}
}

func (u *UserApp) SaveUser(userData entity.User) (user *entity.User, err error) {
	return
}

func (u *UserApp) GetUser(userId uint64) (user *entity.User, err error) {
	return
}

func (u *UserApp) GetUsers() (users []entity.User, err error) {
	return
}

func (u *UserApp) GetUserByEmail(email string) (user *entity.User, err error) {
	return
}
