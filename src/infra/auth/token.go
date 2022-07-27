package auth

import "github.com/gin-gonic/gin"

type TokenDetails struct {
	AccessToken string `json:"access_token"`
	ID          uint64 `json:"id"`
}

type AccessDetail struct {
	UserID uint64
}

type TokenInterface interface {
	CreateToken(userID uint64) (token *TokenDetails, err error)
	ExtractToken(c *gin.Context) (access_detail *AccessDetail, err error)
}

type Token struct {
}

func NewToken() TokenInterface {
	return &Token{}
}

func (t *Token) CreateToken(userID uint64) (token *TokenDetails, err error) {
	return
}

func (t *Token) ExtractToken(c *gin.Context) (access_detail *AccessDetail, err error) {

	access_detail = &AccessDetail{
		UserID: 1,
	}
	return
}
