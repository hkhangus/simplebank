package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HasPassword return the bcrypt hash of the password
func HashPassword(password string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", fmt.Errorf("failed to hash password: %s", password)
	}
	return string(hashedPass), nil
}
// CheckPassword checks if the provided password is correct or not
func CheckPassword(hashedPass, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(password))
}