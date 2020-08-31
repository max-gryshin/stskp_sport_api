package models

import (
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	State     int8      `json:"state"`
	CreatedAt time.Time `json:"created_at"`
	Email     string    `json:"email"`
}
