package middlewares

import (
	"clean-arch-2/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
	// "fmt"
)

type Middleware struct{}

func (m *Middleware) AuthMiddleware() gin.HandlerFunc {
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

		c.Set("userId", tokenClaims["userId"])
		c.Set("roleId", tokenClaims["roleId"])
		c.Next()
		return
	}
}

func (m *Middleware) RoleMiddleware(allowedRoles []uint64) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleId, _ := c.Get("roleId")
		isContains := false

		for _, allowedRole := range allowedRoles {
			if allowedRole == uint64(roleId.(float64)) {
				isContains = true
				break
			}
		}

		if !isContains {
			c.JSON(
				http.StatusUnauthorized,
				utilities.ApiResponse(
					"Anda tidak memiliki akses ke halaman ini",
					false,
					nil,
				),
			)
			c.Abort()
			return
		}

		c.Next()
		return
	}
}
