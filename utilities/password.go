package utilities

import (
	"golang.org/x/crypto/bcrypt"
	"os"
	"strconv"
)

func HashPassword(password string) (string, error) {
	saltRound, err := strconv.Atoi(os.Getenv("SALT_ROUND"))

	if err != nil {
		return "", err
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), saltRound)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
