package repository

import "github.com/andrersp/financeiro-go/src/domain/entity"

type UserRepository interface {
	SaveUser(userData entity.User) (user *entity.User, err error)
	GetUser(userID uint64) (user *entity.User, err error)
	GetUsers() (users []entity.PublicUser, err error)
	GetUserByEmail(email string) (user *entity.User, err error)
}
