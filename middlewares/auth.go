package middlewares

import (
	"clean-arch-2/utilities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthMiddleware struct{}

func (m *AuthMiddleware) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		if header == "" {
			c.JSON(
				http.StatusUnauthorized,
				utilities.ApiResponse(
					"Anda harus login terlebih dahulu",
					false,
					nil,
				),
			)
			c.Abort()
			return
		}
		header = header[len("Bearer "):]
		tokenClaims, err := utilities.DecodeToken(header)
		if err != nil {
			c.JSON(
				http.StatusUnauthorized,
				utilities.ApiResponse(
					"Token tidak valid",
					false,
					nil,
				),
			)
			c.Abort()
			return
		}

		c.Set("userId", tokenClaims["id"])
		c.Set("roleId", tokenClaims["roleId"])
		c.Next()
		return
	}
}
