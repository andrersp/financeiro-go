package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/andrersp/financeiro-go/src/domain/entity"
	"github.com/andrersp/financeiro-go/src/domain/repository"
	"github.com/andrersp/financeiro-go/src/infra/auth"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type userApi struct {
	repository repository.UserRepository
	token      auth.TokenInterface
}

type UserAppInterface interface {
	SaveUser(c *gin.Context)
	GetUser(c *gin.Context)
	GetUsers(c *gin.Context)
	// GetUserByEmail(c *gin.Context)
}

func NewUserApi(repository repository.UserRepository, token auth.TokenInterface) UserAppInterface {
	return &userApi{
		repository: repository,
		token:      token,
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

	// userID, err := strconv.ParseUint(c.Param("userID"), 10, 64)

	access_detail, err := u.token.ExtractTokenAcessDetail(c)
	userID := *&access_detail.UserID

	token, err := u.token.CreateToken(1)
	if err != nil {
		return
	}
	fmt.Println(*token)

	if err != nil {
		return
	}
	user, err := u.repository.GetUser(userID)

	if err != nil {

		if errors.Is(gorm.ErrRecordNotFound, err) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"msg": "User not found",
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": "internal error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": user.PublicUser(),
	})

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
