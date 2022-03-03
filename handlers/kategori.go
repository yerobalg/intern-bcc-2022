package handlers

import (
	"clean-arch-2/config"
	"clean-arch-2/kategori"
	"clean-arch-2/middlewares"
	"clean-arch-2/utilities"
	// "fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type KategoriHandler struct {
	handler    config.Router
	service    kategori.KategoriService
	middleware middlewares.Middleware
}

func (h KategoriHandler) Setup() {
	api := h.handler.BaseRouter
	// api = h.handler.BaseRouter.Use(h.middleware.RoleMiddleware([]uint64{2, 3}))
	{
		api.POST(
			"/kategori",
			h.middleware.AuthMiddleware(),
			h.middleware.RoleMiddleware([]uint64{1}),
			h.TambahKategori,
		)
		api.GET(
			"/kategori",
			h.GetSemuaKategori,
		)
		api.PUT(
			"/kategori/:idKategori",
			h.middleware.AuthMiddleware(),
			h.middleware.RoleMiddleware([]uint64{1}),
			h.UbahKategori,
		)
		api.DELETE(
			"/kategori/:idKategori",
			h.middleware.AuthMiddleware(),
			h.middleware.RoleMiddleware([]uint64{1}),
			h.HapusKategori,
		)
	}
}

func NewKategoriHandler(
	handler config.Router,
	service kategori.KategoriService,
	middleware middlewares.Middleware,
) KategoriHandler {
	return KategoriHandler{
		handler:    handler,
		service:    service,
		middleware: middleware,
	}
}

func (h *KategoriHandler) TambahKategori(c *gin.Context) {
	var input kategori.KategoriInput
	if err := c.BindJSON(&input); err != nil {
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

	kategori := kategori.Kategori{ Nama: input.Nama }

	if err := h.service.Save(&kategori); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utilities.ApiResponse(
				"Terjadi kesalahan sistem",
				false,
				err.Error(),
			),
		)
		return
	}

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Kategori berhasil ditambahkan",
		"data":    kategori,
	})
}

func (h *KategoriHandler) UbahKategori(c *gin.Context) {
	idRaw, _ := c.Params.Get("idAlamat")
	id, _ := strconv.ParseUint(idRaw, 10, 64)

	var input kategori.KategoriInput
	if err := c.BindJSON(&input); err != nil {
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

	res, err := h.service.GetById(id)

	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			c.JSON(
				http.StatusNotFound,
				utilities.ApiResponse(
					"Kategori tidak ditemukan",
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

	res.Nama = input.Nama

	if err := h.service.Save(&res); err != nil {
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
			"Kategori berhasil diubah",
			true,
			res,
		),
	)
}