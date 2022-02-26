package handlers

import (
	// "clean-arch-2/app/formatter"
	"clean-arch-2/app/models"
	"clean-arch-2/app/services"
	"clean-arch-2/config"
	"clean-arch-2/utilities"
	"fmt"

	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	//"strconv"
	// "strings"
)

type DaerahHandler struct {
	handler config.Router
	service services.DaerahService
}

func (daerah DaerahHandler) Setup() {
	api := daerah.handler.BaseRouter
	{
		api.GET("/daerah/provinsi", daerah.GetSemuaProvinsi)
		api.GET("/daerah/kabupaten/:idProvinsi", daerah.GetDaerah("kabupaten"))
		api.GET("/daerah/kecamatan/:idKabupaten", daerah.GetDaerah("kecamatan"))
		api.GET("/daerah/kelurahan/:idKecamatan", daerah.GetDaerah("kelurahan"))
	}
}

func NewDaerahHandler(
	handler config.Router,
	service services.DaerahService,
) DaerahHandler {
	return DaerahHandler{handler: handler, service: service}
}

func (daerah DaerahHandler) GetSemuaProvinsi(c *gin.Context) {

}

func (daerah DaerahHandler) GetDaerah(paramDaerah string) func(*gin.Context) {
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

	return func(c *gin.Context) {
		idProvinsi, _ := c.Params.Get("idProvinsi")
		body := []models.OutputDaerah{}

		if err := daerah.service.GetDaerah(
			&body,
			idProvinsi,
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
}
