package messages

import "github.com/gin-gonic/gin"
// import "clean-arch-2/app/models"

func PrintError(err error, message string) *gin.H {
	return &gin.H{
		"success": false,
		"message": message,
		"error":   err.Error(),
	}
}