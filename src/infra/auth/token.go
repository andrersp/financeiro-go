package auth

import (
	"github.com/gin-gonic/gin"
)

type TokenDetails struct {
	AccessToken string `json:"access_token"`
	ID          uint64 `json:"id"`
}

type AccessDetail struct {
	UserID uint64
}

type TokenInterface interface {
	CreateToken(userID uint64) (token *TokenDetails, err error)
	ExtractTokenAcessDetail(c *gin.Context) (access_detail *AccessDetail, err error)
}

type Token struct {
}

func NewToken() TokenInterface {
	return &Token{}
}

func (t *Token) CreateToken(userID uint64) (token *TokenDetails, err error) {
	// token.ID = userID

	// atClaims := jwt.MapClaims{}
	// atClaims["authorized"] = true
	// atClaims["user_id"] = userID

	// accessToen := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	// token.AccessToken, err = accessToen.SignedString("fazendo teste")

	return
}

func (t *Token) ExtractTokenAcessDetail(c *gin.Context) (access_detail *AccessDetail, err error) {

	access_detail = &AccessDetail{
		UserID: 1,
	}
	return
}
