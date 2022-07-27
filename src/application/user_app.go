package application

import (
	"github.com/andrersp/financeiro-go/src/domain/repository"
	"github.com/gin-gonic/gin"
)

type userApp struct {
	repository repository.UserRepository
}

//UserApp implements the UserAppInterface
// var _ UserAppInterface = &userApp{}

type UserAppInterface interface {
	SaveUser(c *gin.Context)
	GetUser(c *gin.Context)
	GetUsers(c *gin.Context)
	// GetUserByEmail(c *gin.Context)
}

// func (u *userApp) SaveUser(userData entity.User) (user *entity.User, err error) {
// 	return u.us.SaveUser(userData)

// }

// func (u *userApp) GetUser(userId uint64) (user *entity.User, err error) {

// 	fmt.Println("Aqui ")
// 	return
// }

// func (u *userApp) GetUsers() (users []entity.User, err error) {
// 	return
// }

// func (u *userApp) GetUserByEmail(email string) (user *entity.User, err error) {
// 	return
// }
