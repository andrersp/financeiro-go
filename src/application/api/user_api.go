package api

import (
	"errors"
	"net/http"

	"github.com/andrersp/financeiro-go/src/domain/entity"
	"github.com/andrersp/financeiro-go/src/domain/repository"
	"github.com/andrersp/financeiro-go/src/infra/auth"
	"github.com/andrersp/financeiro-go/src/interfaces/response"

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
	UpdateUser(c *gin.Context)
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
		response.Error(c, http.StatusUnprocessableEntity, err)
		return

	}

	// Validate User
	if err := user.Validate(""); err != nil {
		response.Error(c, http.StatusBadRequest, err)
		return
	}

	userEmail, err := u.repository.GetUserByEmail(user.Email)

	if userEmail != nil {
		response.Error(c, http.StatusBadRequest, errors.New("Email exists"))
		return
	}

	newUser, err := u.repository.SaveUser(user)

	if err != nil {
		response.Error(c, http.StatusInternalServerError, errors.New("Internal server error"))
		return

	}

	response.Success(c, http.StatusCreated, newUser.PublicUser())

}

func (u *userApi) GetUser(c *gin.Context) {

	// userID, err := strconv.ParseUint(c.Param("userID"), 10, 64)

	access_detail, err := u.token.ExtractTokenAcessDetail(c.Request)
	userID := access_detail.UserID

	user, err := u.repository.GetUser(userID)

	if err != nil {

		if errors.Is(gorm.ErrRecordNotFound, err) {
			response.Error(c, http.StatusNotFound, errors.New("user not found"))

			return
		}

		response.Error(c, http.StatusInternalServerError, err)
		return
	}

	response.Success(c, http.StatusOK, user.PublicUser())

}

func (u *userApi) GetUsers(c *gin.Context) {

	users, err := u.repository.GetUsers()

	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}

	response.Success(c, http.StatusOK, users)

}

func (u *userApi) UpdateUser(c *gin.Context) {

	access_detail, err := u.token.ExtractTokenAcessDetail(c.Request)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, nil)
		return
	}

	userID := access_detail.UserID

	var userData entity.User

	if err := c.ShouldBindJSON(&userData); err != nil {
		response.Error(c, http.StatusUnprocessableEntity, err)
		return
	}

	if err := userData.Validate("update"); err != nil {
		response.Error(c, http.StatusBadRequest, err)
		return
	}

	// Verify Email Duplicate

	userEmail, err := u.repository.GetUserByEmail(userData.Email)

	if err != nil && !errors.Is(gorm.ErrRecordNotFound, err) {
		response.Error(c, http.StatusBadRequest, err)
		return
	}

	if userEmail != nil {
		if userEmail.ID != userID {
			response.Error(c, http.StatusBadRequest, errors.New("Email exist!"))
			return
		}
	}

	// Select User By Email
	user, err := u.repository.GetUser(userID)
	if user == nil {
		response.Error(c, http.StatusNotFound, errors.New("User Not Found"))
		return
	}

	// Set New Data
	user.Email = userData.Email
	user.FirstName = userData.FirstName
	user.LastName = userData.LastName

	user, err = u.repository.UpdateUser(*user)
	if err != nil {
		return
	}

	response.Success(c, http.StatusOK, user.PublicUser())

}
