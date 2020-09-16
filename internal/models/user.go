package models

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

const (
	StateHalfRegistration = 1
	StateRegistration     = 2
	StateActive           = 3
	StateBlocked          = 4
	StateDeleted          = 5
)

type Users []User

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username" valid:"Required; MaxSize(50)"`
	Password  string    `json:"password" valid:"Required; MaxSize(50)"`
	State     int8      `json:"state"`
	CreatedAt time.Time `json:"created_at"`
	Email     string    `json:"email"`
}

// SetPassword sets a new password stored as hash.
func (u *User) SetPassword(password string) error {
	if bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14); err != nil {
		return err
	} else {
		u.Password = string(bytes)

		return nil
	}
}

// InvalidPassword returns true if the given password does not match the hash.
func (u *User) InvalidPassword(password string) bool {
	if u.Password == "" && password == "" {
		return false
	}

	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	return err != nil
}
