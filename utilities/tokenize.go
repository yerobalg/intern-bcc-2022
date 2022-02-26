package utilities

import (
	"clean-arch-2/app/models"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func EncodeToken(user *models.Users) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":     user.ID,
		"roleId": user.RoleID,
		"exp":    time.Now().Add(time.Hour * 24 * 7).Unix(),
	}).SignedString([]byte("secret"))
}
