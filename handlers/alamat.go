package handlers

import (
	"clean-arch-2/alamat"
	"clean-arch-2/config"
	"clean-arch-2/middlewares"
	"clean-arch-2/role"
	"clean-arch-2/utilities"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type AlamatHandler struct {
	handler     config.Router
	service     alamat.AlamatService
	roleService role.RoleService
	middleware  middlewares.Middleware
}

func (h AlamatHandler) Setup() {
	api := h.handler.BaseRouter
	// api = h.handler.BaseRouter.Use(h.middleware.RoleMiddleware([]uint64{2, 3}))
	{
		api.POST(
			"/alamat",
			h.middleware.AuthMiddleware(),
			h.middleware.RoleMiddleware([]uint64{2, 3}),
			h.TambahAlamat,
		)
		api.GET("/alamat/:idAlamat")
		api.PUT(
			"/alamat/:idAlamat",
			h.middleware.AuthMiddleware(),
			h.UbahAlamat,
		)
		api.DELETE("/alamat/:idAlamat")
	}
}

func NewAlamatHandler(
	handler config.Router,
	service alamat.AlamatService,
	roleService role.RoleService,
	middleware middlewares.Middleware,
) AlamatHandler {
	return AlamatHandler{
		handler:     handler,
		service:     service,
		roleService: roleService,
		middleware:  middleware,
	}
}

func (h *AlamatHandler) TambahAlamat(c *gin.Context) {
	userIdInterface, _ := c.Get("userId")
	roleIdInterface, _ := c.Get("roleId")
	userId := uint64(userIdInterface.(float64))
	roleId := uint64(roleIdInterface.(float64))
	var body alamat.AlamatInput

	if err := c.BindJSON(&body); err != nil {
		c.JSON(
			http.StatusBadRequest,
			utilities.ApiResponse(
				"Validasi Error",
				false,
				utilities.FormatBindError(err),
			),
		)
		return
	}

	alamatObj := alamat.Alamat{
		Nama:          body.Nama,
		AlamatLengkap: body.AlamatLengkap,
		KodePos:       body.KodePos,
		IDKelurahan:   body.IDKelurahan,
		IDUser:        userId,
		IsUser:        roleId == 2,
	}

	if err := h.service.Save(&alamatObj); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utilities.ApiResponse(
				"Terjadi kesalahan Sistem",
				false,
				err.Error(),
			),
		)
		return
	}

	role, _ := h.roleService.GetRoleById(roleId)

	c.JSON(
		http.StatusOK,
		utilities.ApiResponse(
			"Alamat "+role.Nama+" berhasil ditambahkan",
			true,
			alamat.AlamatInputFormatter(&alamatObj),
		),
	)
}

func (h *AlamatHandler) UbahAlamat(c *gin.Context) {
	idRaw, _ := c.Params.Get("idAlamat")
	id, _ := strconv.ParseUint(idRaw, 10, 64)

	userIdInterface, _ := c.Get("userId")
	userId := uint64(userIdInterface.(float64))

	var body alamat.AlamatInput
	if err := c.BindJSON(&body); err != nil {
		c.JSON(
			http.StatusBadRequest,
			utilities.ApiResponse(
				"Validasi Error",
				false,
				utilities.FormatBindError(err),
			),
		)
		return
	}

	res, err := h.service.GetById(id)

	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			c.JSON(
				http.StatusNotFound,
				utilities.ApiResponse(
					"Alamat tidak ditemukan",
					false,
					nil,
				),
			)
		} else {
			fmt.Println(err)
			c.JSON(
				http.StatusInternalServerError,
				utilities.ApiResponse(
					"Terjadi kesalahan Sistem",
					false,
					err.Error(),
				),
			)
		}
		return
	}

	if (res.IDUser != userId) {
		c.JSON(
			http.StatusForbidden,
			utilities.ApiResponse(
				"Anda tidak memiliki akses untuk mengubah alamat ini",
				false,
				nil,
			),
		)
		return
	}

	res.Nama = body.Nama
	res.AlamatLengkap = body.AlamatLengkap
	res.KodePos = body.KodePos
	res.IDKelurahan = body.IDKelurahan

	if err := h.service.Update(res); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utilities.ApiResponse(
				"Terjadi kesalahan Sistem",
				false,
				err.Error(),
			),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		utilities.ApiResponse(
			"Alamat berhasil diubah",
			true,
			alamat.AlamatInputFormatter(res),
		),
	)
}
