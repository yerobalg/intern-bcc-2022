package handlers

import (
	"clean-arch-2/app/services"
	"clean-arch-2/app/models"
	"clean-arch-2/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	handler    config.Router
	service services.PostService
}

func (post PostHandler) Setup() {
	api := post.handler.BaseRouter
	{
		api.GET("/post", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "ping pong",
			})
		})
		api.GET("/post/:id")
		api.POST("/post", post.Save)
		api.POST("/post/:id")
		api.DELETE("/post/:id")
	}

}

func NewPostHandler(
	handler config.Router,
	service services.PostService,
) PostHandler {
	return PostHandler{handler: handler, service: service}
}

func (post PostHandler) Save(c *gin.Context) {
	var postInput = models.PostInput{}
	if err := c.BindJSON(&postInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	postModel := models.Post{
		Title: postInput.Title,
		Text:  postInput.Text,
	}

	if err := post.service.Save(&postModel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Error saving post",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Post saved successfully",
		"data":    postModel,
	})
}
