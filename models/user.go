package models

import (
	"html"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint32 `json:"id"`
	Nickname  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

}

func VerifyPassword(hashedPassword, password string) error {

	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

}

func (u *User) BeforeSave() error {
	hashedpassword, err := Hash(u.Password)
	if err != nil {
		panic(err)
	}

	u.Password = string(hashedpassword)
	return nil

}

func (u *User) Preapare() {
	u.ID = 0
	u.Nickname = html.EscapeString(strings.TrimSpace(u.Nickname))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

}
