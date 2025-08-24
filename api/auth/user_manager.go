package auth

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/hasona23/workis/api/helpers"
	"golang.org/x/crypto/bcrypt"
)

const (
	ADMIN  = "Admin"
	HELPER = "Helper"
	NONE   = "Un-Determined"
)
const HASH_COST = 16

func AddUser(userRequest UserCreateRequest) error {
	if err := helpers.IsValidEmail(userRequest.Email); err != nil {
		return err
	}
	if err := ValdiatePassword(userRequest.Password); err != nil {
		return err
	}
	passwordHash, err := HashPassword(userRequest.Password)
	if err != nil {
		return err
	}
	user := User{
		ID:           uuid.New(),
		Email:        userRequest.Email,
		PasswordHash: passwordHash,
		Role:         NONE,
	}
	UserDB.Create(user)
	return nil
}
func HashPassword(password string) (string, error) {
	passData := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(passData, HASH_COST)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
func ValdiatePassword(password string) error {

	err := helpers.IsStringInBounds(password, "Password", 16, 256)
	if err != nil {
		return err
	}
	if strings.ToLower(password) == password || strings.ToUpper(password) == password {
		return fmt.Errorf("password must have upper and lowercase letter")
	}

	if !strings.ContainsAny(password, "0123456789") {
		return fmt.Errorf("password must contain a number")
	}
	if !strings.ContainsAny(strings.ToLower(password), "abcdefghijklmnopqrstuvwxyz") {
		return fmt.Errorf("password must contain letters")
	}
	if !strings.ContainsAny(password, "+-_@£$%*&#><?.,") {
		return fmt.Errorf("password must contain a special character +-_@£$%%*&#><?.,")
	}
	if strings.ContainsAny(password, "[]{}\" !()}~|/\\") {
		return fmt.Errorf("password cannot contain: []{}\" !()}~|/\\")
	}
	return nil
}
