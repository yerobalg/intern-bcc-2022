package handlers

import (
	"clean-arch-2/config"
	"clean-arch-2/kategori"
	"clean-arch-2/middlewares"
	"clean-arch-2/produk"
	"clean-arch-2/utilities"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	// "io"
	//"log"
	"os"
	"net/http"
	"strconv"
	"strings"
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
			h.middleware.RoleMiddleware([]uint64{1}),
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
			h.middleware.RoleMiddleware([]uint64{1}),
			h.UbahProduk,
		)
		api.DELETE(
			"/produk/:slug",
			h.middleware.AuthMiddleware(),
			h.middleware.RoleMiddleware([]uint64{1}),
			h.HapusProduk,
		)
		api.GET(
			"/produk/:slug",
			h.GetProdukBySlug,
		)
		api.POST(
			"produk/gambar/:slug",
			h.middleware.AuthMiddleware(),
			h.middleware.RoleMiddleware([]uint64{1}),
			h.TambahGambarProduk,
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

	if len(body.IdTags) == 0 {
		c.JSON(
			http.StatusBadRequest,
			utilities.ApiResponse(
				"Tags tidak boleh kosong",
				false,
				nil,
			),
		)
		return
	}

	produkObj := produk.Produk{
		NamaProduk: body.NamaProduk,
		Slug:       slug.Make(body.NamaProduk),
		Harga:      body.Harga,
		Diskon:     body.Diskon,
		Stok:       body.Stok,
		Deskripsi:  body.Deskripsi,
		IsHiasan:   body.IsHiasan,
		Gender:     body.Gender,
		Berat:      body.Berat,
	}

	idTags := body.IdTags
	idTags2 := append(idTags, body.IDKategori)

	if err := h.service.Save(&produkObj, idTags2); err != nil {
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
			produk.ProdukInputFormatter(&produkObj, body.IDKategori, idTags),
		),
	)
}

func (h *ProdukHandler) UbahProduk(c *gin.Context) {
	slug, _ := c.Params.Get("slug")
	fmt.Println(slug)
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
	res.Diskon = body.Diskon
	res.Stok = body.Stok
	res.Deskripsi = body.Deskripsi
	res.IsHiasan = body.IsHiasan
	res.Gender = body.Gender
	res.Berat = body.Berat

	idTags := body.IdTags
	idTags2 := append(idTags, body.IDKategori)

	if err := h.service.Update(res, idTags2); err != nil {
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
			produk.ProdukInputFormatter(res, body.IDKategori, idTags),
		),
	)
}

func (h *ProdukHandler) GetProdukBySlug(c *gin.Context) {
	slug, _ := c.Params.Get("slug")

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

	c.JSON(
		http.StatusOK,
		utilities.ApiResponse(
			"Berhasil mengambil produk",
			true,
			produk.ProdukOutputFormatter(*res),
		),
	)
}

func (h *ProdukHandler) HapusProduk(c *gin.Context) {
	slug, _ := c.Params.Get("slug")

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

	if err := h.service.Delete(res); err != nil {
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
			"Produk berhasil dihapus",
			true,
			produk.ProdukOutputFormatter(*res),
		),
	)
}

func (h *ProdukHandler) TambahGambarProduk(c *gin.Context) {
	slug, _ := c.Params.Get("slug")
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

	//Multiple Form
	form, err := c.MultipartForm()
	if err != nil {
		panic(err)
	}
	files := form.File["gambar"]

	var url []string

	// For range
	for i, file := range files {
		path := fmt.Sprintf(
			"public/images/products/%s_%s.%s",
			strconv.Itoa(i+1),
			slug,
			strings.Split(file.Filename, ".")[1],
		)
		if err := c.SaveUploadedFile(file, path); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("err: %s", err.Error()))
			return
		}
		url = append(url, path)
	}

	err = h.service.SaveGambar(uint64(res.ID), url)
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

	for i := 0; i < len(url); i++ {
		url[i] = os.Getenv("BASE_URL") + "/" + url[i]
	}

	// Response
	c.JSON(
		http.StatusOK,
		utilities.ApiResponse(
			"Berhasil Mengupload Gambar", 
			true, 
			gin.H{"gambar": url},
		),
	)
}


