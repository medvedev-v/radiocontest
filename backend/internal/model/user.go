package model

import (
	"time"
)

type User struct {
	ID           int    `json:"id"`
	UserName         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
	UserRole         string `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}
