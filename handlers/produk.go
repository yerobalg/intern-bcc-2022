package handlers

import (
	"clean-arch-2/config"
	"clean-arch-2/kategori"
	"clean-arch-2/middlewares"
	"clean-arch-2/produk"
	"clean-arch-2/utilities"
	"fmt"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"

	// "io"
	//"log"
	"net/http"
	"os"
	// "path"
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
	{
		api.POST(
			"/produk",
			h.middleware.AuthMiddleware(),
			h.middleware.RoleMiddleware([]uint64{1}),
			h.TambahProduk,
		)
		api.GET(
			"/produk/search/",
			h.CariProduk,
		)
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
		api.DELETE(
			"produk/gambar/:slug/",
			h.middleware.AuthMiddleware(),
			h.middleware.RoleMiddleware([]uint64{1}),
			h.HapusGambarProduk,
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

	if len(body.IdTags) == 0 && body.IsHiasan == false {
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

	produkObj, err := h.service.GetBySlug(slug)
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

	idKategori := produkObj.KategoriProduk[len(produkObj.KategoriProduk)-1].IDKategori

	hiasanProduk, err, hasil := h.service.GetHiasanProduk(idKategori)

	var idHiasan []uint64
	for _, h := range hiasanProduk {
		idHiasan = append(idHiasan, h.ID)
	}

	gambarHiasan, err := h.service.GetGambarProduk(idHiasan)

	c.JSON(
		http.StatusOK,
		utilities.ApiResponse(
			"Berhasil mengambil produk",
			true,
			produk.ProdukOutputFormatter(
				*produkObj,
				produk.ProdukSearchFormatter(hiasanProduk, gambarHiasan, hasil),
			),
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
			nil),
	)
}

func (h *ProdukHandler) CariProduk(c *gin.Context) {
	tagsKey, isTagsExist := c.GetQuery("tags")
	kategoriKey, isKategoriExist := c.GetQuery("kategori")
	kataKunciKey, _ := c.GetQuery("kata-kunci")
	maxHargaKey, isMaxHargaExist := c.GetQuery("max-harga")
	pageKey, isPageExist := c.GetQuery("page")
	genderKey, isGenderExist := c.GetQuery("gender")
	sortKey, isSortExist := c.GetQuery("sort")

	if !isGenderExist {
		genderKey = "BOTH"
	} else {
		genderKey = strings.ToUpper(genderKey)
	}

	pageNum := 1
	if isPageExist {
		pageInt, _ := strconv.Atoi(pageKey)
		pageNum = pageInt
	}

	if !isSortExist {
		sortKey = "popular"
	}

	maxHarga := uint64(0)
	if isMaxHargaExist {
		maxHargaInt, _ := strconv.Atoi(maxHargaKey)
		maxHarga = uint64(maxHargaInt)
	}

	var kategoriList []uint64
	if isKategoriExist {
		kategoriInt, _ := strconv.Atoi(kategoriKey)
		kategoriList = append(kategoriList, uint64(kategoriInt))
	}

	if isTagsExist {
		tags := strings.Split(tagsKey, ",")
		for _, tag := range tags {
			tagInt, _ := strconv.Atoi(tag)
			kategoriList = append(kategoriList, uint64(tagInt))
		}
	}

	searchRes, err, jumlah := h.service.CariProduk(
		kategoriList,
		kataKunciKey,
		maxHarga,
		genderKey,
		pageNum,
		sortKey,
	)
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

	for _, produk := range searchRes {
		idProduk = append(idProduk, produk.ID)
	}

	res, err := h.service.GetGambarProduk(idProduk)
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
			"Berhasil mengambil produk",
			true,
			produk.ProdukSearchFormatter(searchRes, res, jumlah),
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

func (h *ProdukHandler) HapusGambarProduk(c *gin.Context) {
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

	for _, gambar := range res.GambarProduk {
		dir, _ := filepath.Abs("./" + gambar.Nama)
		if err := os.Remove(dir); err != nil {
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
	}

	err = h.service.DeleteGambarProduk(uint64(res.ID))
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
			"Berhasil Menghapus Gambar",
			true,
			nil,
		),
	)

}
