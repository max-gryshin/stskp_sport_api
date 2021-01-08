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

var userFields = map[string][]string{
	"get":    {"id", "user_name", "state", "created_at", "email"},
	"update": {"user_name", "state", "email"},
}

type User struct {
	ID        int       `json:"id" db:"id" binding:"required"`
	Username  string    `json:"username"   valid:"MaxSize(50)" db:"user_name"`
	Password  string    `json:"password"   valid:"MaxSize(50)" db:"password_hash"`
	State     int8      `json:"state"      valid:"Range(1, 5)"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	Email     string    `json:"email"      valid:"Email"`
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

func GetAllowedUserFieldsByMethod(method string) []string {
	return userFields[method]
}
