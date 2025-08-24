package auth

import (
	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID
	Email        string `gorm:"unique"`
	PasswordHash string
	Role         string
}

type UserCreateRequest struct {
	Email    string
	Password string
}
