package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hash), err
}

func ValidatePassword(hashpashword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashpashword), []byte(password))
	return err
}
