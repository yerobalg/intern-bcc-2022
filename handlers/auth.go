package handlers

import (
	"clean-arch-2/user"
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
	service user.UserService
}

func (h AuthHandler) Setup() {
	api := h.handler.BaseRouter
	{
		api.POST("/auth/login", h.Login)
		api.POST("/auth/register", h.Register)
	}
}

func NewAuthHandler(
	handler config.Router,
	service user.UserService,
) AuthHandler {
	return AuthHandler{handler: handler, service: service}
}

func (h AuthHandler) Login(c *gin.Context) {
	var body user.UserLoginInput

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

	userObj, err := h.service.Login(&body)
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

	if err := utilities.CheckPassword(body.Password, userObj.Password); err != nil {
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

	token, err := utilities.EncodeToken(userObj)
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
			user.LoginFormat(userObj, token),
		),
	)
}

//Register handler
func (h AuthHandler) Register(c *gin.Context) {
	var body user.UserRegisterInput

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

	userObj := user.Users{
		Username:     body.Username,
		Email:        body.Email,
		Password:     password,
		JenisKelamin: body.JenisKelamin,
		NomorHp:      body.NomorHp,
		Nama:         body.Nama,
		RoleID:       body.RoleID,
	}

	if err := h.service.Register(&userObj); err != nil {
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
	if (userObj.RoleID == 2) {
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
			user.RegisterFormat(&userObj),
		),
	)
}
