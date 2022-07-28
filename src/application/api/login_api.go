package api

import (
	"errors"
	"net/http"

	"github.com/andrersp/financeiro-go/src/domain/entity"
	"github.com/andrersp/financeiro-go/src/domain/repository"
	"github.com/andrersp/financeiro-go/src/http/response"
	"github.com/andrersp/financeiro-go/src/infra/auth"
	"github.com/andrersp/financeiro-go/src/infra/security"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type loginApi struct {
	repository repository.UserRepository
	token      auth.TokenInterface
}

type LoginApiInterface interface {
	Login(c *gin.Context)
	// GetUserByEmail(c *gin.Context)
}

func NewLoginApi(userRepo repository.UserRepository, token auth.TokenInterface) LoginApiInterface {

	return &loginApi{
		repository: userRepo,
		token:      token,
	}

}

func (l *loginApi) Login(c *gin.Context) {

	var user *entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.Error(c, http.StatusUnprocessableEntity, err)
		return
	}
	if err := user.Validate("login"); err != nil {
		response.Error(c, http.StatusUnprocessableEntity, err)
		return
	}

	u, err := l.repository.GetUserByEmail(user.Email)

	if err != nil {

		if errors.Is(gorm.ErrRecordNotFound, err) {
			response.Error(c, http.StatusUnprocessableEntity, errors.New("Err Email or Password"))
			return
		}

		response.Error(c, http.StatusInternalServerError, err)
		return
	}

	if err := security.CheckPassword(u.Password, user.Password); err != nil {
		response.Error(c, http.StatusUnprocessableEntity, errors.New("Err Email or Password"))
		return
	}

	tokenData, err := l.token.CreateToken(u.ID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}

	response.Success(c, http.StatusOK, tokenData)

}
