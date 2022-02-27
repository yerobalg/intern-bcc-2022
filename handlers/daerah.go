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
	middleware middlewares.AuthMiddleware
}

func (h DaerahHandler) Setup() {
	api := h.handler.BaseRouter.Use(h.middleware.AuthMiddleware())
	{
		api.GET("/daerah/provinsi", h.GetSemuaProvinsi)
		api.GET("/daerah/kabupaten/:idProvinsi", h.GetDaerah)
		api.GET("/daerah/kecamatan/:idKabupaten", h.GetDaerah)
		api.GET("/daerah/kelurahan/:idKecamatan", h.GetDaerah)
	}
}

func NewDaerahHandler(
	handler config.Router,
	service daerah.DaerahService,
	middleware middlewares.AuthMiddleware,
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
				err.Error(),
				http.StatusInternalServerError,
				"Gagal mendapatkan seluruh provinsi",
				nil,
			),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		utilities.ApiResponse(
			"Berhasil mendapatkan seluruh provinsi",
			http.StatusOK,
			"Sukses",
			&body,
		),
	)
}

func (h DaerahHandler) GetDaerah(c *gin.Context) {
	split := strings.Split(c.Request.URL.Path, "/")
	paramDaerah := split[4]
	paramID := split[5]

	var daerahAsal string
	var daerahTujuan string

	switch paramDaerah {
	case "kabupaten":
		daerahAsal = "kab"
		daerahTujuan = "prov"
	case "kecamatan":
		daerahAsal = "kec"
		daerahTujuan = "kab"
	case "kelurahan":
		daerahAsal = "kel"
		daerahTujuan = "kec"
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
				http.StatusInternalServerError,
				"Gagal mendapatkan daerah",
				nil,
			),
		)
		return
	}

	if len(body) == 0 {
		c.JSON(
			http.StatusNotFound,
			utilities.ApiResponse(
				fmt.Sprintf("%s tidak ditemukan", paramDaerah),
				http.StatusNotFound,
				"Data tidak ditemukan",
				nil,
			),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		utilities.ApiResponse(
			fmt.Sprintf("Berhasil mengambil %s", paramDaerah),
			http.StatusOK,
			"Sukses",
			&body,
		),
	)
}
