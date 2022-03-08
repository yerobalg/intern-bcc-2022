package handlers

import (
	"clean-arch-2/alamat"
	"clean-arch-2/config"
	"clean-arch-2/keranjang"
	"clean-arch-2/middlewares"
	"clean-arch-2/produk"
	"clean-arch-2/user"
	"clean-arch-2/utilities"
	//"fmt"

	// "fmt"
	"net/http"
	// "strconv"
	// "strings"

	"github.com/gin-gonic/gin"
)

type KeranjangHandler struct {
	handler       config.Router
	service       keranjang.KeranjangService
	userService   user.UserService
	alamatService alamat.AlamatService
	produkService produk.ProdukService
	middleware    middlewares.Middleware
}

func (h KeranjangHandler) Setup() {
	api := h.handler.BaseRouter
	{
		api.POST(
			"/keranjang",
			h.middleware.AuthMiddleware(),
			h.middleware.RoleMiddleware([]uint64{2}),
			h.TambahKeranjang,
		)
		api.GET(
			"/keranjang",
			h.middleware.AuthMiddleware(),
			h.middleware.RoleMiddleware([]uint64{2}),
		)
	}

}

func NewKeranjangHandler(
	handler config.Router,
	service keranjang.KeranjangService,
	userService user.UserService,
	alamatService alamat.AlamatService,
	produkService produk.ProdukService,
	middleware middlewares.Middleware,
) KeranjangHandler {
	return KeranjangHandler{
		handler:       handler,
		service:       service,
		userService:   userService,
		alamatService: alamatService,
		produkService: produkService,
		middleware:    middleware,
	}
}

func (h *KeranjangHandler) TambahKeranjang(c *gin.Context) {
	var body keranjang.InputKeranjang

	userIdInterface, _ := c.Get("userId")
	userId := uint64(userIdInterface.(float64))

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

	res, _ := h.produkService.GetBySlug(body.Slug)

	if res.Stok < uint(body.JumlahBeli) {
		c.JSON(
			http.StatusBadRequest,
			utilities.ApiResponse(
				"Stok tidak mencukupi",
				false,
				nil,
			),
		)
		return
	}

	keranjangSeller := keranjang.Keranjang{
		IDUser:   userId,
	}

	err := h.service.AddKeranjang(&keranjangSeller)

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

	err = h.service.AddKeranjangProduk(&keranjang.Keranjang_Produk{
		IDKeranjang: uint64(keranjangSeller.ID),
		IDProduk:    uint64(res.ID),
		Jumlah:      body.JumlahBeli,
	})

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
			"Berhasil menambahkan keranjang",
			true,
			keranjangSeller.ID,
		),
	)
}
