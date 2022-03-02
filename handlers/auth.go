package handlers

import (
	"clean-arch-2/config"
	"clean-arch-2/user"
	"clean-arch-2/role"
	"clean-arch-2/utilities"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	handler config.Router
	service user.UserService
	roleService role.RoleService
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
	roleService role.RoleService,
) AuthHandler {
	return AuthHandler{
		handler: handler, 
		service: service, 
		roleService: roleService,
	}
}

func (h AuthHandler) Login(c *gin.Context) {
	var body user.UserLoginInput

	if err := c.BindJSON(&body); err != nil {
		c.JSON(
			http.StatusBadRequest,
			utilities.ApiResponse(
				"Harap isi semua data",
				false,
				utilities.FormatBindError(err),
			),
		)
		return
	}

	fieldContainSpaces := utilities.FieldContainSpaces([]utilities.Field{
		{
			Name:  "username atau email",
			Value: body.UsernameOrEmail,
		},
		{
			Name:  "password",
			Value: body.Password,
		},
	})

	if len(fieldContainSpaces) > 0 {
		c.JSON(
			http.StatusBadRequest,
			utilities.ApiResponse(
				"Terdapat field yang tidak boleh mengandung spasi",
				false,
				gin.H{"errors": fieldContainSpaces},
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
					false,
					nil,
				),
			)
		} else {
			c.JSON(
				http.StatusInternalServerError,
				utilities.ApiResponse(
					"Gagal melakukan login",
					false,
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
				false,
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
				false,
				nil,
			),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		utilities.ApiResponse(
			"Login Sukses",
			true,
			user.LoginFormat(userObj, token),
		),
	)
}

//Register handler
func (h AuthHandler) Register(c *gin.Context) {
	var body user.UserRegisterInput

	if err := c.BindJSON(&body); err != nil {
		c.JSON(
			http.StatusBadRequest,
			utilities.ApiResponse(
				"Validasi error",
				false,
				utilities.FormatBindError(err),
			),
		)
		return
	}

	fieldContainSpaces := utilities.FieldContainSpaces([]utilities.Field{
		{
			Name:  "password",
			Value: body.Password,
		},
	})

	if len(fieldContainSpaces) > 0 {
		c.JSON(
			http.StatusBadRequest,
			utilities.ApiResponse(
				"Terdapat field yang tidak boleh mengandung spasi",
				false,
				gin.H{"errors": fieldContainSpaces},
			),
		)
		return
	}

	if !utilities.IsValidUsername(body.Username) {
		c.JSON(
			http.StatusBadRequest,
			utilities.ApiResponse(
				"Username hanya boleh mengandung huruf, angka, titik dan underscore",
				false,
				nil,
			),
		)
		return
	}

	if !utilities.IsValidPassword(body.Password) {
		c.JSON(
			http.StatusBadRequest,
			utilities.ApiResponse(
				"Password harus mengandung minimal satu huruf dan satu angka",
				false,
				nil,
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
				false,
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
					"Username, email, atau nomor HP sudah terdaftar",
					false,
					nil,
				),
			)
		} else {
			c.JSON(
				http.StatusInternalServerError,
				utilities.ApiResponse(
					"Gagal membuat user",
					false,
					nil,
				),
			)
		}
		return
	}

	role, _ := h.roleService.GetRoleById(userObj.RoleID)

	c.JSON(
		http.StatusCreated,
		utilities.ApiResponse(
			fmt.Sprintf("Registrasi %s berhasil", role.Nama),
			true,
			user.RegisterFormat(&userObj),
		),
	)
}
