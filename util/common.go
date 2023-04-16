package util

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func Attempt(stmt error) {
	if stmt != nil {
		log.Fatal("fatal", stmt)
	}
}

// Returns computes hash from given password string
func GeneratePasswordHash(password string) (string, error) {
	if hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost); err != nil {
		return "", err
	} else {
		return string(hash), nil
	}
}

// Returns bool value by comparing passwordHash with given original password
func ComparePasswordHash(passwordHash string, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password)) == nil
}
