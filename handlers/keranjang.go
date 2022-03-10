package handlers

import (
	"clean-arch-2/alamat"
	"clean-arch-2/config"
	"clean-arch-2/keranjang"
	"clean-arch-2/middlewares"
	"clean-arch-2/produk"
	"clean-arch-2/user"
	"clean-arch-2/utilities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
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
			h.GetKeranjangUser,
		)
		api.POST(
			"/keranjang/konfirmasi",
			h.middleware.AuthMiddleware(),
			h.middleware.RoleMiddleware([]uint64{2}),
			h.KonfirmasiKeranjang,
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
		IDProduk: uint64(res.ID),
		Jumlah:   uint(body.JumlahBeli),
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
			gin.H{"idKeranjang": keranjangSeller.ID},
		),
	)
}

func (h *KeranjangHandler) GetKeranjangUser(c *gin.Context) {
	userIdInterface, _ := c.Get("userId")
	userId := uint64(userIdInterface.(float64))

	keranjangObj, err := h.service.GetKeranjangUser(userId)

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

	alamatUser, err := h.alamatService.GetAllUserAddress(userId)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			c.JSON(
				http.StatusBadRequest,
				utilities.ApiResponse(
					"User belum memiliki alamat",
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
			"Berhasil mengambil keranjang",
			true,
			keranjang.KeranjangFormatter(
				&keranjangObj,
				alamat.GetUserAlamatFormatter(&alamatUser),
			),
		),
	)
}

func (h *KeranjangHandler) KonfirmasiKeranjang(c *gin.Context) {
	var body keranjang.InputKonfirmasi

	userIdInterface, _ := c.Get("userId")
	userId := uint64(userIdInterface.(float64))

	c.BindJSON(&body)

	keranjangObj, err := h.service.GetKeranjangBatch(body.IDKeranjang)
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

	formatKeranjang := keranjang.KeranjangFormatter(&keranjangObj, nil)

	userObj, _ := h.userService.GetByID(userId)

	alamatUser, err := h.alamatService.GetById(body.IDAlamat)
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

	alamatAdmin, err := h.alamatService.GetAdminAddress()
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

	metodePembayaran, err := h.service.GetSemuaMetodeBayar()

	ongkosKirim := utilities.CekOngkosKirim(
		alamatUser.IDKabupaten,
		alamatAdmin.IDKabupaten,
		int(formatKeranjang.TotalBerat),
	)

	c.JSON(
		http.StatusOK,
		utilities.ApiResponse(
			"Berhasil menuju menu konfirmasi",
			true,
			keranjang.KonfirmasiPesananFormatter(
				user.ProfilUserFormatter(userObj),
				&keranjangObj,
				alamat.GetAlamatByIdFormatter(*alamatUser),
				alamat.GetAlamatByIdFormatter(alamatAdmin),
				ongkosKirim,
				metodePembayaran,
			),
		),
	)
}
