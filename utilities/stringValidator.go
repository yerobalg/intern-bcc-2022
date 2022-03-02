package utilities

import (
	"unicode"
)

func IsValidPassword(password string) bool {
	isLetter := false
	isDigit := false
	for _, char := range password {
		if unicode.IsLetter(char) {
			isLetter = true
		}
		if unicode.IsNumber(char) {
			isDigit = true
		}
		if isLetter && isDigit {
			return true
		}
	}
	return false
}

func IsValidUsername(username string) bool {
	for _, char := range username {
		if 
			!unicode.IsLetter(char) && !unicode.IsNumber(char) && 
			char != '_' && char != '.' {
			return false
		}
	}
	return true
}