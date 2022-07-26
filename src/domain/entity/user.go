package entity

import (
	"html"
	"strings"
	"time"

	"github.com/andrersp/financeiro-go/src/infra/security"
	"gorm.io/gorm"
)

type User struct {
	ID        uint64    `json:"id" `
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
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