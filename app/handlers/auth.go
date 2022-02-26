package handlers

import (
	"clean-arch-2/app/models"
	"clean-arch-2/app/services"
	"clean-arch-2/config"
	"clean-arch-2/utilities"
	"clean-arch-2/utilities/messages"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AuthHandler struct {
	handler config.Router
	service services.UserService
}

func (auth AuthHandler) Setup() {
	api := auth.handler.BaseRouter
	{
		api.POST("/auth/login", auth.Login)
		api.GET("/auth/test", func(c *gin.Context) {
			c.JSON(
				http.StatusOK,
				gin.H{
					"message": "Hello World",
				},
			)
		})
		api.POST("/auth/register", auth.Register)
	}
}

func NewAuthHandler(
	handler config.Router,
	service services.UserService,
) AuthHandler {
	return AuthHandler{handler: handler, service: service}
}

func (auth AuthHandler) Login(c *gin.Context) {

}

//Register handler
func (auth AuthHandler) Register(c *gin.Context) {
	var body models.UserRegisterInput

	if err := c.BindJSON(&body); err != nil {
		c.JSON(
			http.StatusBadRequest,
			messages.PrintError(err, "Invalid request body"),
		)
		return
	}

	password, err := utilities.HashPassword(body.Password)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			messages.PrintError(err, "Failed to hash password"),
		)
	}

	user := models.Users{
		Username:     body.Username,
		Email:        body.Email,
		Password:     password,
		JenisKelamin: body.JenisKelamin,
		NomorHp:      body.NomorHp,
		Nama:         body.Nama,
		RoleID:       2,
	}

	if err := auth.service.Register(&user); err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			c.JSON(
				http.StatusBadRequest,
				messages.PrintError(err, "Username or email already exist"),
			)
		} else {
			c.JSON(
				http.StatusInternalServerError,
				messages.PrintError(err, "Failed to register user to database"),
			)
		}
		return
	}

	c.JSON(
		http.StatusOK,
		messages.RegisterSuccess(&user),
	)
}
