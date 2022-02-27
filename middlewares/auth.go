package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"clean-arch-2/utilities"
)

type AuthMiddleware struct {}

func (m *AuthMiddleware) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		if header == "" {
			c.JSON(
				http.StatusUnauthorized,
				utilities.ApiResponse(
					"You must login to get access",
					http.StatusUnauthorized,
					"Unauthorized",
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
					http.StatusUnauthorized,
					"Token tidak valid",
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