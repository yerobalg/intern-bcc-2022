package handlers

import (
	"clean-arch-2/alamat"
	"clean-arch-2/config"
	"clean-arch-2/middlewares"
	"clean-arch-2/role"
	"clean-arch-2/utilities"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
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
			h.middleware.RoleMiddleware([]uint64{1, 2}),
			h.TambahAlamat,
		)
		api.GET(
			"/alamat",
			h.middleware.AuthMiddleware(),
			h.GetAlamatUser,
		)
		api.GET(
			"/alamat/:idAlamat",
			h.middleware.AuthMiddleware(),
			h.GetAlamatById,
		)
		api.PUT(
			"/alamat/:idAlamat",
			h.middleware.AuthMiddleware(),
			h.UbahAlamat,
		)
		api.DELETE(
			"/alamat/:idAlamat",
			h.middleware.AuthMiddleware(),
			h.middleware.RoleMiddleware([]uint64{2}),
			h.HapusAlamat,
		)
		api.GET(
			"/alamat/admin",
			h.GetAlamatAdmin,
		)
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

	if roleId == 1 {
		res, err := h.service.GetAllUserAddress(userId)
		if err != nil {
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

		fmt.Println("tes")

		if len(res) > 0 {
			c.JSON(
				http.StatusBadRequest,
				utilities.ApiResponse(
					"Admin sudah memiliki alamat",
					false,
					nil,
				),
			)
			return
		}
	}

	alamatObj := alamat.Alamat{
		Nama:          body.Nama,
		AlamatLengkap: body.AlamatLengkap,
		KodePos:       body.KodePos,
		IDKabupaten:   body.IDKabupaten,
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

	if res.IDUser != userId {
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
	res.IDKabupaten = body.IDKabupaten

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
			alamat.GetAlamatByIdFormatter(*res),
		),
	)
}

func (h *AlamatHandler) HapusAlamat(c *gin.Context) {
	idRaw, _ := c.Params.Get("idAlamat")
	id, _ := strconv.ParseUint(idRaw, 10, 64)

	userIdInterface, _ := c.Get("userId")
	userId := uint64(userIdInterface.(float64))

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

	if res.IDUser != userId {
		c.JSON(
			http.StatusForbidden,
			utilities.ApiResponse(
				"Anda tidak memiliki akses untuk menghapus alamat ini",
				false,
				nil,
			),
		)
		return
	}

	if err := h.service.Delete(res); err != nil {
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
			"Alamat berhasil dihapus",
			true,
			alamat.GetAlamatByIdFormatter(*res),
		),
	)
}

func (h *AlamatHandler) GetAlamatUser(c *gin.Context) {
	userIdInterface, _ := c.Get("userId")
	userId := uint64(userIdInterface.(float64))

	res, err := h.service.GetAllUserAddress(userId)

	if err != nil {
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
			"Alamat berhasil diambil",
			true,
			alamat.GetUserAlamatFormatter(res),
		),
	)
}

func (h *AlamatHandler) GetAlamatById(c *gin.Context) {
	idRaw, _ := c.Params.Get("idAlamat")
	id, _ := strconv.ParseUint(idRaw, 10, 64)

	userIdInterface, _ := c.Get("userId")
	userId := uint64(userIdInterface.(float64))

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

	if res.IDUser != userId {
		c.JSON(
			http.StatusForbidden,
			utilities.ApiResponse(
				"Anda tidak memiliki akses untuk melihat alamat ini",
				false,
				nil,
			),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		utilities.ApiResponse(
			"Alamat berhasil diambil",
			true,
			alamat.GetAlamatByIdFormatter(*res),
		),
	)
}

func (h *AlamatHandler) GetAlamatAdmin(c *gin.Context) {
	res, err := h.service.GetAdminAddress()

	if err != nil {
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
			"Alamat Admin berhasil diambil",
			true,
			alamat.GetAlamatByIdFormatter(res),
		),
	)
}
