package handlers

import (
	"clean-arch-2/config"
	"clean-arch-2/kategori"
	"clean-arch-2/middlewares"
	"clean-arch-2/produk"
	"clean-arch-2/utilities"
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

type ProdukHandler struct {
	handler         config.Router
	service         produk.ProdukService
	KategoriService kategori.KategoriService
	middleware      middlewares.Middleware
}

func (h ProdukHandler) Setup() {
	api := h.handler.BaseRouter
	// api = h.handler.BaseRouter.Use(h.middleware.RoleMiddleware([]uint64{2, 3}))
	{
		api.POST(
			"/produk",
			h.middleware.AuthMiddleware(),
			h.middleware.RoleMiddleware([]uint64{3}),
			h.TambahProduk,
		)
		api.GET(
			"/produk",
			// h.GetAlamatUser,
		)
		// api.GET(
		// 	"/alamat/:idAlamat",
		// 	h.middleware.AuthMiddleware(),
		// 	h.GetAlamatById,
		// )
		api.PUT(
			"/produk/:slug",
			h.middleware.AuthMiddleware(),
			// h.UbahAlamat,
		)
		api.DELETE(
			"/produk/:slug",
			h.middleware.AuthMiddleware(),
			h.middleware.RoleMiddleware([]uint64{2}),
			// h.HapusAlamat,
		)
		api.GET(
			"/produk/:slug",
			// h.GetAlamatSeller,
		)
	}
}

func NewProdukHandler(
	handler config.Router,
	service produk.ProdukService,
	kategoriService kategori.KategoriService,
	middleware middlewares.Middleware,
) ProdukHandler {
	return ProdukHandler{
		handler:         handler,
		service:         service,
		KategoriService: kategoriService,
		middleware:      middleware,
	}
}

func (h *ProdukHandler) TambahProduk(c *gin.Context) {
	userIdInterface, _ := c.Get("userId")
	userId := uint64(userIdInterface.(float64))

	var body produk.ProdukInput
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

	produkObj := produk.Produk{
		NamaProduk: body.NamaProduk,
		Slug:       slug.Make(body.NamaProduk),
		Harga:      body.Harga,
		Stok:       body.Stok,
		Deskripsi:  body.Deskripsi,
		IDSeller:   userId,
		IsHiasan:   body.IsHiasan,
		Gender:     body.Gender,
	}

	if err := h.service.Save(&produkObj); err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			c.JSON(
				http.StatusBadRequest,
				utilities.ApiResponse(
					"Produk sudah ada",
					false,
					err.Error(),
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
			"Produk berhasil ditambahkan",
			true,
			produk.GetProdukFormat(&produkObj),
		),
	)
}

func (h *ProdukHandler) UbahProduk(c *gin.Context) {
	slug, _ := c.Params.Get("slug")
	var body produk.ProdukInput
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

	res, err := h.service.GetBySlug(slug)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			c.JSON(
				http.StatusNotFound,
				utilities.ApiResponse(
					"Produk tidak ditemukan",
					false,
					err.Error(),
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

	res.NamaProduk = body.NamaProduk
	res.Harga = body.Harga
	res.Stok = body.Stok
	res.Deskripsi = body.Deskripsi
	res.IsHiasan = body.IsHiasan
	res.Gender = body.Gender

	if err := h.service.Update(res); err != nil {
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
			"Produk berhasil diubah",
			true,
			produk.GetProdukFormat(res),
		),
	)
}
