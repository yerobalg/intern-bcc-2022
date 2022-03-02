package handlers

import (
	"clean-arch-2/config"
	"clean-arch-2/daerah"
	"clean-arch-2/middlewares"
	"clean-arch-2/utilities"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
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
			h.GetDaerah,
		)
		api.GET("/daerah/kecamatan/:idKabupaten",
			h.middleware.AuthMiddleware(),
			h.GetDaerah,
		)
		api.GET(
			"/daerah/kelurahan/:idKecamatan",
			h.middleware.AuthMiddleware(),
			h.GetDaerah,
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
			&body,
		),
	)
}

func (h DaerahHandler) GetDaerah(c *gin.Context) {
	split := strings.Split(c.Request.URL.Path, "/")
	paramDaerah := split[4]
	paramID := split[5]

	var daerahAsal, daerahTujuan, paramTujuan string

	switch paramDaerah {
	case "kabupaten":
		daerahAsal = "kab"
		daerahTujuan = "prov"
		paramTujuan = "ID provinsi"
	case "kecamatan":
		daerahAsal = "kec"
		daerahTujuan = "kab"
		paramTujuan = "ID Kabupaten"
	case "kelurahan":
		daerahAsal = "kel"
		daerahTujuan = "kec"
		paramTujuan = "ID Kecamatan"
	}

	daerahAsal = "id_" + daerahAsal
	daerahTujuan = "id_" + daerahTujuan

	body := []daerah.OutputDaerah{}

	if err := h.service.GetDaerah(
		&body,
		paramID,
		daerahAsal,
		daerahTujuan,
		paramDaerah,
	); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utilities.ApiResponse(
				err.Error(),
				false,
				nil,
			),
		)
		return
	}

	if len(body) == 0 {
		c.JSON(
			http.StatusNotFound,
			utilities.ApiResponse(
				fmt.Sprintf("%s tidak ditemukan", paramTujuan),
				false,
				nil,
			),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		utilities.ApiResponse(
			fmt.Sprintf("Berhasil mengambil %s", paramDaerah),
			true,
			&body,
		),
	)
}
