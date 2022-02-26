package handlers

import (
	"clean-arch-2/app/formatter"
	"clean-arch-2/app/models"
	"clean-arch-2/app/services"
	"clean-arch-2/config"
	"clean-arch-2/utilities"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	//"strconv"
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
	var body models.UserLoginInput

	if err := c.BindJSON(&body); err != nil {
		var message string
		
		c.JSON(
			http.StatusBadRequest,
			utilities.ApiResponse(
				message,
				http.StatusBadRequest,
				"Validation Error",
				utilities.FormatBindError(err),
			),
		)
		return
	}

	user, err := auth.service.Login(&body)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			c.JSON(
				http.StatusBadRequest,
				utilities.ApiResponse(
					"Username atau email tidak terdaftar",
					http.StatusBadRequest,
					"Validation Error",
					nil,
				),
			)
		} else {
			c.JSON(
				http.StatusInternalServerError,
				utilities.ApiResponse(
					"Gagal melakukan login",
					http.StatusInternalServerError,
					"Internal Server Error",
					nil,
				),
			)
		}
		return
	}

	if err := utilities.CheckPassword(body.Password, user.Password); err != nil {
		c.JSON(
			http.StatusBadRequest,
			utilities.ApiResponse(
				"Password salah",
				http.StatusBadRequest,
				"Validation Error",
				nil,
			),
		)
		return
	}

	token, err := utilities.EncodeToken(user)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utilities.ApiResponse(
				"Gagal membuat token",
				http.StatusInternalServerError,
				"Internal Server Error",
				nil,
			),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		utilities.ApiResponse(
			"Login Sukses",
			http.StatusOK,
			"Sukses",
			formatter.LoginUser(user, token),
		),
	)
}

//Register handler
func (auth AuthHandler) Register(c *gin.Context) {
	var body models.UserRegisterInput

	if err := c.BindJSON(&body); err != nil {
		var message string
		if (!strings.Contains(err.Error(), "'min' tag")) {
			message = "Harap isi semua data"
		} else {
			message = "Panjang password minimal 6 karakter"
		}
		c.JSON(
			http.StatusBadRequest,
			utilities.ApiResponse(
				message,
				http.StatusBadRequest,
				"Validation Error",
				utilities.FormatBindError(err),
			),
		)
		return
	}

	password, err := utilities.HashPassword(body.Password)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utilities.ApiResponse(
				"Gagal membuat password",
				http.StatusInternalServerError,
				"Internal Server Error",
				nil,
			),
		)
	}

	user := models.Users{
		Username:     body.Username,
		Email:        body.Email,
		Password:     password,
		JenisKelamin: body.JenisKelamin,
		NomorHp:      body.NomorHp,
		Nama:         body.Nama,
		RoleID:       body.RoleID,
	}

	if err := auth.service.Register(&user); err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			c.JSON(
				http.StatusBadRequest,
				utilities.ApiResponse(
					"Username atau email sudah terdaftar",
					http.StatusBadRequest,
					"Validation Error",
					nil,
				),
			)
		} else {
			c.JSON(
				http.StatusInternalServerError,
				utilities.ApiResponse(
					"Gagal membuat user",
					http.StatusInternalServerError,
					"Internal Server Error",
					nil,
				),
			)
		}
		return
	}

	var role string
	if (user.RoleID == 2) {
		role = "User"
	} else {
		role = "Seller"
	}

	c.JSON(
		http.StatusOK,
		utilities.ApiResponse(
			fmt.Sprintf("Registrasi %s berhasil", role),
			http.StatusOK,
			"Sukses",
			formatter.RegisterUser(&user),
		),
	)
}
