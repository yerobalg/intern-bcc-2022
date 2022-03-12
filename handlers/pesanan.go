package handlers

import (
	"clean-arch-2/alamat"
	"clean-arch-2/config"
	"clean-arch-2/keranjang"
	"clean-arch-2/middlewares"
	"clean-arch-2/pesanan"
	"clean-arch-2/produk"
	"clean-arch-2/utilities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
	// "fmt"
)

type PesananHandler struct {
	handler          config.Router
	service          pesanan.PesananService
	keranjangService keranjang.KeranjangService
	alamatService    alamat.AlamatService
	produkService    produk.ProdukService
	middleware       middlewares.Middleware
}

func (h PesananHandler) Setup() {
	api := h.handler.BaseRouter
	{
		api.POST(
			"/pesanan",
			h.middleware.AuthMiddleware(),
			h.middleware.RoleMiddleware([]uint64{2}),
			h.TambahPesanan,
		)
		api.GET(
			"/pesanan/:idPesanan",
			h.middleware.AuthMiddleware(),
			h.middleware.RoleMiddleware([]uint64{2}),
			h.GetPesananByID,
		)
		api.PATCH(
			"/pesanan/:idPesanan/dibayar",
			h.middleware.AuthMiddleware(),
			h.middleware.RoleMiddleware([]uint64{2}),
			h.UpdateStatusPesanan("diproses"),
		)
		api.PATCH(
			"/pesanan/:idPesanan/dikirim",
			h.middleware.AuthMiddleware(),
			h.middleware.RoleMiddleware([]uint64{1}),
			h.UpdateStatusPesanan("dikirim"),
		)
	}
}

func NewPesananHandler(
	handler config.Router,
	service pesanan.PesananService,
	keranjangService keranjang.KeranjangService,
	alamatService alamat.AlamatService,
	produkService produk.ProdukService,
	middleware middlewares.Middleware,
) PesananHandler {
	return PesananHandler{
		handler:          handler,
		service:          service,
		keranjangService: keranjangService,
		alamatService:    alamatService,
		produkService:    produkService,
		middleware:       middleware,
	}
}

func (h *PesananHandler) TambahPesanan(c *gin.Context) {

	userIdInterface, _ := c.Get("userId")
	userId := uint64(userIdInterface.(float64))

	var body pesanan.InputPesanan
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

	res, err := h.keranjangService.GetMetodeByID(body.IDMetodePembayaran)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			c.JSON(
				http.StatusBadRequest,
				utilities.ApiResponse(
					"Metode Pembayaran tidak ditemukan",
					false,
					err.Error(),
				),
			)
		} else {
			c.JSON(
				http.StatusInternalServerError,
				utilities.ApiResponse(
					"Terjadi kesalahan sistem",
					false,
					err.Error(),
				),
			)
		}
		return
	}

	keranjangObj, err := h.keranjangService.GetKeranjangBatch(body.IDKeranjang)
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

	var idProduk []uint64
	var idKeranjang []uint64
	var jumlah []uint

	for _, r := range keranjangObj {
		if r.IsPaid {
			c.JSON(
				http.StatusBadRequest,
				utilities.ApiResponse(
					"Keranjang sudah diproses",
					false,
					nil,
				),
			)
			return
		}
		idProduk = append(idProduk, r.IDProduk)
		jumlah = append(jumlah, r.Jumlah)
		idKeranjang = append(idKeranjang, uint64(r.ID))
	}

	pesananObj := pesanan.Pesanan{
		KodeKurir:        body.KodeKurir,
		KodeLayanan:      body.KodeLayanan,
		HargaOngkosKirim: body.HargaOngkosKirim,
		EstimasiKirim:    body.EstimasiKirim,
		BatasBayar:       time.Unix(time.Now().Unix()+(int64(res.ExpiredTime)*60), 0),
		IDUser:           userId,
		IDMetode:         body.IDMetodePembayaran,
		IDAlamat:         body.IDAlamat,
	}

	if err = h.service.AddPesanan(&pesananObj, body.IDKeranjang); err != nil {
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

	if err = h.produkService.KurangiStok(idProduk, jumlah); err != nil {
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

	if err = h.keranjangService.KeranjangDipesan(idKeranjang); err != nil {
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
		http.StatusCreated,
		utilities.ApiResponse(
			"Pesanan berhasil ditambahkan",
			true,
			pesanan.PesananInputFormatter(&pesananObj, body.IDKeranjang),
		),
	)
}

func (h *PesananHandler) GetPesananByID(c *gin.Context) {
	id, _ := c.Params.Get("idPesanan")
	idPesanan, _ := strconv.ParseUint(id, 10, 64)

	userIdInterface, _ := c.Get("userId")
	userId := uint64(userIdInterface.(float64))

	roleIdInterface, _ := c.Get("roleId")
	roleId := uint64(roleIdInterface.(float64))

	pesananObj, err := h.service.GetPesananByID(idPesanan)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			c.JSON(
				http.StatusNotFound,
				utilities.ApiResponse(
					"Pesanan tidak ditemukan",
					false,
					err.Error(),
				),
			)
		} else {
			c.JSON(
				http.StatusInternalServerError,
				utilities.ApiResponse(
					"Terjadi kesalahan sistem",
					false,
					err.Error(),
				),
			)
		}
		return
	}

	if pesananObj.IDUser != userId && roleId != 1 {
		c.JSON(
			http.StatusForbidden,
			utilities.ApiResponse(
				"Anda tidak memiliki akses untuk pesanan ini",
				false,
				nil,
			),
		)
		return
	}

	alamatAdmin, err := h.alamatService.GetAdminAddress()
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

	res, err := h.service.GetPesananByID(idPesanan)
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
			"Pesanan berhasil diambil",
			true,
			pesanan.PesananOutputFormatter(&res, alamatAdmin),
		),
	)
}

func (h *PesananHandler) UpdateStatusPesanan(status string) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, _ := c.Params.Get("idPesanan")
		idPesanan, _ := strconv.ParseUint(id, 10, 64)

		userIdInterface, _ := c.Get("userId")
		userId := uint64(userIdInterface.(float64))

		roleIdInterface, _ := c.Get("roleId")
		roleId := uint64(roleIdInterface.(float64))

		pesananObj, err := h.service.GetPesananByID(idPesanan)
		if err != nil {
			if strings.Contains(err.Error(), "record not found") {
				c.JSON(
					http.StatusNotFound,
					utilities.ApiResponse(
						"Pesanan tidak ditemukan",
						false,
						err.Error(),
					),
				)
			} else {
				c.JSON(
					http.StatusInternalServerError,
					utilities.ApiResponse(
						"Terjadi kesalahan sistem",
						false,
						err.Error(),
					),
				)
			}
			return
		}

		if pesananObj.IDUser != userId && roleId != 1 {
			c.JSON(
				http.StatusForbidden,
				utilities.ApiResponse(
					"Anda tidak memiliki akses ke pesanan ini",
					false,
					nil,
				),
			)
			return
		}

		if 
			pesananObj.BatasBayar.Unix() < time.Now().UTC().Add(time.Hour * 7).Unix() && 
			status == "diproses" {
			c.JSON(
				http.StatusForbidden,
				utilities.ApiResponse(
					"Batas pembayaran telah habis",
					false,
					nil,
				),
			)
			return
		}

		if !pesananObj.TanggalBayar.IsZero() && status == "diproses" {
			c.JSON(
				http.StatusForbidden,
				utilities.ApiResponse(
					"Pesanan telah dibayar",
					false,
					nil,
				),
			)
			return
		}

		if pesananObj.TanggalBayar.IsZero() && status == "dikirim" {
			c.JSON(
				http.StatusForbidden,
				utilities.ApiResponse(
					"Pesanan belum dibayar",
					false,
					nil,
				),
			)
			return
		}

		if err = h.service.UpdatePesanan(&pesananObj, status); err != nil {
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
				"Status pesanan berhasil diubah menjadi "+status,
				true,
				nil,
			),
		)
	}
}
