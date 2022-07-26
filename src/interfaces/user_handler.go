package interfaces

import (
	"net/http"

	"github.com/andrersp/financeiro-go/src/application"
	"github.com/andrersp/financeiro-go/src/domain/entity"
	"github.com/gin-gonic/gin"
)

type User struct {
	userApp application.UserAppInterface
}

func NewUserHandler(userAppInterface application.UserAppInterface) *User {
	return &User{
		userApp: userAppInterface,
	}
}
func (u *User) SaveUser(c *gin.Context) {

	var user entity.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			gin.H{
				"error": err.Error(),
			},
		)
		return

	}

	newUser, err := u.userApp.SaveUser(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal error",
		})
		return

	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    newUser.PublicUser(),
	})

}

func (u *User) GetUser(c *gin.Context) {

}

func (u *User) GetUsers(c *gin.Context) {

}
