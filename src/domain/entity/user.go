package entity

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/andrersp/financeiro-go/src/infra/security"
	"github.com/badoux/checkmail"
	"gorm.io/gorm"
)

type User struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id" `
	FirstName string    `gorm:"size:100;not null;" json:"first_name"`
	LastName  string    `gorm:"size:100;not null;" json:"last_name"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}

type PublicUser struct {
	ID        uint64 `json:"id"`
	FirstName string `json:"first_name"`
	Email     string `json:"email"`
}

func (u *User) Prepare() {
	u.FirstName = html.EscapeString(strings.TrimSpace(u.FirstName))
	u.LastName = html.EscapeString(strings.TrimSpace(u.LastName))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	hash, err := security.HashGenerator(u.Password)
	if err != nil {
		return
	}

	u.Password = string(hash)

	return
}

//So that we dont expose the user's email address and password to the world
func (u *User) PublicUser() interface{} {
	return &PublicUser{
		ID:        u.ID,
		FirstName: u.FirstName,
		Email:     u.Email,
	}
}

func (u *User) Validate(action string) (err error) {

	if u.Email == "" {
		err = errors.New("email is required")
		return

	}

	if err = checkmail.ValidateFormat(u.Email); err != nil {
		err = errors.New("please provide a valid email")
		return

	}

	switch strings.ToLower(action) {
	case "update":
		if u.FirstName == "" {
			err = errors.New("first name is required")
			return
		}
		if u.LastName == "" {
			err = errors.New("last name is required")
			return
		}

	case "login":
		if u.Password == "" {
			err = errors.New("password is required")
			return
		}
	case "forgotpassword":
		return
	default:
		if u.FirstName == "" {
			err = errors.New("first name is required")
			return
		}
		if u.LastName == "" {
			err = errors.New("last name is required")
			return
		}
		if u.Password == "" {
			err = errors.New("password is required")
			return
		}
		if u.Password != "" && len(u.Password) < 6 {
			err = errors.New("password should be at least 6 characters")
			return
		}

	}
	return
}
