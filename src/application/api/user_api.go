package api

import (
	"net/http"

	"github.com/andrersp/financeiro-go/src/domain/entity"
	"github.com/andrersp/financeiro-go/src/domain/repository"
	"github.com/gin-gonic/gin"
)

type userApi struct {
	repository repository.UserRepository
}

type UserAppInterface interface {
	SaveUser(c *gin.Context)
	GetUser(c *gin.Context)
	GetUsers(c *gin.Context)
	// GetUserByEmail(c *gin.Context)
}

func NewUserApi(repository repository.UserRepository) UserAppInterface {
	return &userApi{
		repository: repository,
	}
}

func (u *userApi) SaveUser(c *gin.Context) {

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

	// Validate User
	if err := user.Validate(""); err != nil {
		c.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			gin.H{
				"error": err.Error(),
			},
		)
		return

	}

	newUser, err := u.repository.SaveUser(user)

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

func (u *userApi) GetUser(c *gin.Context) {

	users, err := u.repository.GetUsers()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(200, users)

}

func (u *userApi) GetUsers(c *gin.Context) {

	users, err := u.repository.GetUsers()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(200, users)

}
