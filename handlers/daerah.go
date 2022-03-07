package handlers

import (
	"clean-arch-2/config"
	"clean-arch-2/daerah"
	"clean-arch-2/middlewares"
	"clean-arch-2/utilities"
	"github.com/gin-gonic/gin"
	"net/http"
	//"strings"
)

type DaerahHandler struct {
	handler    config.Router
	service    daerah.DaerahService
	middleware middlewares.Middleware
}

func (h DaerahHandler) Setup() {
	api := h.handler.BaseRouter
	{
		api.GET(
			"/daerah/provinsi",
			h.middleware.AuthMiddleware(),
			h.GetSemuaProvinsi,
		)
		api.GET(
			"/daerah/kabupaten/:idProvinsi",
			h.middleware.AuthMiddleware(),
			h.GetKabupaten,
		)
	}
}

func NewDaerahHandler(
	handler config.Router,
	service daerah.DaerahService,
	middleware middlewares.Middleware,
) DaerahHandler {
	return DaerahHandler{
		handler:    handler,
		service:    service,
		middleware: middleware,
	}
}

func (h DaerahHandler) GetSemuaProvinsi(c *gin.Context) {
	body := []daerah.Provinsi{}

	if err := h.service.GetSemuaProvinsi(&body); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utilities.ApiResponse(
				"Kesalahan sistem gagal mendapatkan seluruh provinsi",
				false,
				nil,
			),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		utilities.ApiResponse(
			"Berhasil mendapatkan seluruh provinsi",
			true,
			daerah.ProvinsiFormat(&body),
		),
	)
}

func (h DaerahHandler) GetKabupaten(c *gin.Context) {
	idProvinsi := c.Param("idProvinsi")

	res, err := h.service.GetKabupaten(idProvinsi)
	if err != nil {
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

	if len(res) == 0 {
		c.JSON(
			http.StatusNotFound,
			utilities.ApiResponse(
				"ID Provinsi tidak ditemukan",
				false,
				nil,
			),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		utilities.ApiResponse(
			"Berhasil mendapatkan seluruh kabupaten",
			true,
			daerah.KabupatenFormat(&res),
		),
	)
}
