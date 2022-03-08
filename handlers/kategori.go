package handlers

import (
	"clean-arch-2/config"
	"clean-arch-2/kategori"
	"clean-arch-2/middlewares"
	"clean-arch-2/utilities"
	//"fmt"

	// "fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
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
			h.GetKategoriOrTag(false),
		)
		api.GET(
			"/kategori/:idKategori",
			h.GetKategoriByID,
		)
		api.GET(
			"/kategori/tag",
			h.GetKategoriOrTag(true),
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

func (h *KategoriHandler) GetKategoriOrTag(isTag bool) func(c *gin.Context) {
	return func(c *gin.Context) {
		res, err := h.service.GetSemuaKategori(isTag)
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

		c.JSON(
			http.StatusOK,
			utilities.ApiResponse(
				"Berhasil mengambil semua "+isTagName(isTag),
				true,
				kategori.GetSemuaKategoriFormatter(&res),
			),
		)
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

	kategoriObj := kategori.Kategori{Nama: input.Nama, IsTag: input.IsTag}

	if err := h.service.Save(&kategoriObj); err != nil {
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

	c.JSON(
		http.StatusOK,
		utilities.ApiResponse(
			"Berhasil menambahkan "+isTagName(kategoriObj.IsTag),
			true,
			kategori.GetKategoriFormatter(&kategoriObj),
		),
	)
}

func (h *KategoriHandler) UbahKategori(c *gin.Context) {
	idRaw, _ := c.Params.Get("idKategori")
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
					isTagName(input.IsTag)+" tidak ditemukan",
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

	if err := h.service.Update(&res); err != nil {
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
			isTagName(res.IsTag)+" berhasil diubah",
			true,
			kategori.GetKategoriFormatter(&res),
		),
	)
}

func (h *KategoriHandler) HapusKategori(c *gin.Context) {
	idRaw, _ := c.Params.Get("idKategori")
	id, _ := strconv.ParseUint(idRaw, 10, 64)

	res, err := h.service.GetById(id)

	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			c.JSON(
				http.StatusNotFound,
				utilities.ApiResponse(
					"Tag atau kategori tidak ditemukan",
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

	if err := h.service.Delete(&res); err != nil {
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
			isTagName(res.IsTag)+" berhasil dihapus",
			true,
			kategori.GetKategoriFormatter(&res),
		),
	)
}

func (h *KategoriHandler) GetKategoriByID(c *gin.Context) {
	idRaw, _ := c.Params.Get("idKategori")
	id, _ := strconv.ParseUint(idRaw, 10, 64)

	res, err := h.service.GetById(id)

	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			c.JSON(
				http.StatusNotFound,
				utilities.ApiResponse(
					"Kategori atau tag tidak ditemukan",
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

	c.JSON(
		http.StatusOK,
		utilities.ApiResponse(
			"Berhasil mengambil "+isTagName(res.IsTag),
			true,
			kategori.GetKategoriFormatter(&res),
		),
	)
}

func isTagName(isTag bool) string {
	if isTag {
		return "tag"
	} else {
		return "kategori"
	}
}
