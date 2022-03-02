package utilities

import (
	"clean-arch-2/user"
	"github.com/golang-jwt/jwt/v4"
	"time"
	"os"
)

func EncodeToken(user *user.Users) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":     user.ID,
		"roleId": user.RoleID,
		"exp":    time.Now().Add(time.Hour * 24 * 7).Unix(),
	}).SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
}

func DecodeToken(token string) (map[string]interface{}, error) {
	decoded, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := decoded.Claims.(jwt.MapClaims)

	if ok && decoded.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
