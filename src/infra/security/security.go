package security

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func HashGenerator(password string) (hash []byte, err error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func CheckPassword(passwordHash, passwordPlain string) (err error) {

	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(passwordPlain))
	if err != nil {
		err = errors.New("Invalid Passowrd")
	}

	return
}
