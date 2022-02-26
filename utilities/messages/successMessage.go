package messages

import "github.com/gin-gonic/gin"
import "clean-arch-2/app/models"

func LoginSuccess(user *models.Users, token string) *gin.H {
	return &gin.H{
		"success": true,
		"message": "User successfully logged in.",
		"data": gin.H{
			"id":    user.ID,
			"token": token,
		},
	}
}

func RegisterSuccess(user *models.Users) *gin.H {
	return &gin.H{
		"success": true,
		"message": "User successfully registered.",
		"data": gin.H{
			"id": user.ID,
			"username": user.Username,
			"email": user.Email,
		},
	}
}
