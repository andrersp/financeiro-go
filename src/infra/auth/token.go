package auth

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
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
	ExtractTokenAcessDetail(r *http.Request) (access_detail *AccessDetail, err error)
}

type Token struct {
}

func NewToken() TokenInterface {

	return &Token{}
}

func (t *Token) CreateToken(userID uint64) (token *TokenDetails, err error) {

	token = &TokenDetails{}

	token.ID = userID
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userID
	atClaims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()

	accessToen := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token.AccessToken, err = accessToen.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return
	}

	return

}

func (t *Token) ExtractTokenAcessDetail(r *http.Request) (access_detail *AccessDetail, err error) {

	token, err := VerifyToken(r)

	var userID uint64

	if err != nil {
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userID, err = strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 64)
		if err != nil {
			return
		}
		access_detail = &AccessDetail{
			UserID: userID,
		}
	}

	return
}

func TokenValid(r *http.Request) (err error) {
	token, err := VerifyToken(r)

	if err != nil {
		return
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return
	}
	return
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {

	tokenString := ExtractToken(r)

	token, err := jwt.Parse(tokenString, verifyTokenKey)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func verifyTokenKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Token not valid! %v", token.Header["alg"])
	}
	return []byte(os.Getenv("SECRET_KEY")), nil
}

func ExtractToken(r *http.Request) (token string) {
	bearToken := r.Header.Get("Authorization")

	splitToken := strings.Split(bearToken, " ")

	if len(splitToken) == 2 {
		token = splitToken[1]
	}
	return
}
