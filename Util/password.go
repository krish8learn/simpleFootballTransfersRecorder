package Util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

//hashing of password
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("hashing failed: %v", err)
	}

	return string(hashedPassword), nil
}

// input input password with provided password
func CheckPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
