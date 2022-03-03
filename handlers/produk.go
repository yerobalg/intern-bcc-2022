package handlers

import (
	"clean-arch-2/produk"
	"clean-arch-2/kategori"
	"clean-arch-2/config"
	"clean-arch-2/middlewares"
	// "clean-arch-2/utilities"
	// "net/http"
	// "strconv"
	// "strings"
	// "github.com/gin-gonic/gin"
)

type ProdukHandler struct {
	handler     config.Router
	service     produk.ProdukService
	KategoriService kategori.KategoriService
	middleware  middlewares.Middleware
}

func (h ProdukHandler) Setup() {
	api := h.handler.BaseRouter
	// api = h.handler.BaseRouter.Use(h.middleware.RoleMiddleware([]uint64{2, 3}))
	{
		api.POST(
			"/produk",
			h.middleware.AuthMiddleware(),
			h.middleware.RoleMiddleware([]uint64{3}),
			// h.TambahProduk,
		)
		api.GET(
			"/produk/:idProduk",
			// h.GetAlamatUser,
		)
		// api.GET(
		// 	"/alamat/:idAlamat",
		// 	h.middleware.AuthMiddleware(),
		// 	h.GetAlamatById,
		// )
		api.PUT(
			"/produk/:idProduk",
			h.middleware.AuthMiddleware(),
			// h.UbahAlamat,
		)
		api.DELETE(
			"/produk/:idProduk",
			h.middleware.AuthMiddleware(),
			h.middleware.RoleMiddleware([]uint64{2}),
			// h.HapusAlamat,
		)
		api.GET(
			"/produk/:idProduk",
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
		handler:     handler,
		service:     service,
		KategoriService: kategoriService,
		middleware:  middleware,
	}
}

// func (h *AlamatHandler) TambahAlamat(c *gin.Context) {
// 	userIdInterface, _ := c.Get("userId")
// 	roleIdInterface, _ := c.Get("roleId")
// 	userId := uint64(userIdInterface.(float64))
// 	roleId := uint64(roleIdInterface.(float64))
// 	var body alamat.AlamatInput

// 	if err := c.BindJSON(&body); err != nil {
// 		c.JSON(
// 			http.StatusBadRequest,
// 			utilities.ApiResponse(
// 				"Validasi Error",
// 				false,
// 				utilities.FormatBindError(err),
// 			),
// 		)
// 		return
// 	}

// 	if roleId == 3 {
// 		res, err := h.service.GetAllUserAddress(userId)
// 		if err != nil {
// 			c.JSON(
// 				http.StatusInternalServerError,
// 				utilities.ApiResponse(
// 					"Terjadi kesalahan Sistem",
// 					false,
// 					err.Error(),
// 				),
// 			)
// 			return
// 		}

// 		if len(res) > 0 {
// 			c.JSON(
// 				http.StatusBadRequest,
// 				utilities.ApiResponse(
// 					"Seller sudah memiliki alamat",
// 					false,
// 					nil,
// 				),
// 			)
// 			return
// 		}
// 	}

// 	alamatObj := alamat.Alamat{
// 		Nama:          body.Nama,
// 		AlamatLengkap: body.AlamatLengkap,
// 		KodePos:       body.KodePos,
// 		IDKelurahan:   body.IDKelurahan,
// 		IDUser:        userId,
// 		IsUser:        roleId == 2,
// 	}

// 	if err := h.service.Save(&alamatObj); err != nil {
// 		c.JSON(
// 			http.StatusInternalServerError,
// 			utilities.ApiResponse(
// 				"Terjadi kesalahan Sistem",
// 				false,
// 				err.Error(),
// 			),
// 		)
// 		return
// 	}

// 	role, _ := h.roleService.GetRoleById(roleId)

// 	c.JSON(
// 		http.StatusOK,
// 		utilities.ApiResponse(
// 			"Alamat "+role.Nama+" berhasil ditambahkan",
// 			true,
// 			alamat.AlamatInputFormatter(&alamatObj),
// 		),
// 	)
// }

// func (h *AlamatHandler) UbahAlamat(c *gin.Context) {
// 	idRaw, _ := c.Params.Get("idAlamat")
// 	id, _ := strconv.ParseUint(idRaw, 10, 64)

// 	userIdInterface, _ := c.Get("userId")
// 	userId := uint64(userIdInterface.(float64))

// 	var body alamat.AlamatInput
// 	if err := c.BindJSON(&body); err != nil {
// 		c.JSON(
// 			http.StatusBadRequest,
// 			utilities.ApiResponse(
// 				"Validasi Error",
// 				false,
// 				utilities.FormatBindError(err),
// 			),
// 		)
// 		return
// 	}

// 	res, err := h.service.GetById(id)

// 	if err != nil {
// 		if strings.Contains(err.Error(), "record not found") {
// 			c.JSON(
// 				http.StatusNotFound,
// 				utilities.ApiResponse(
// 					"Alamat tidak ditemukan",
// 					false,
// 					nil,
// 				),
// 			)
// 		} else {
// 			c.JSON(
// 				http.StatusInternalServerError,
// 				utilities.ApiResponse(
// 					"Terjadi kesalahan Sistem",
// 					false,
// 					err.Error(),
// 				),
// 			)
// 		}
// 		return
// 	}

// 	if res.IDUser != userId {
// 		c.JSON(
// 			http.StatusForbidden,
// 			utilities.ApiResponse(
// 				"Anda tidak memiliki akses untuk mengubah alamat ini",
// 				false,
// 				nil,
// 			),
// 		)
// 		return
// 	}

// 	res.Nama = body.Nama
// 	res.AlamatLengkap = body.AlamatLengkap
// 	res.KodePos = body.KodePos
// 	res.IDKelurahan = body.IDKelurahan

// 	if err := h.service.Update(res); err != nil {
// 		c.JSON(
// 			http.StatusInternalServerError,
// 			utilities.ApiResponse(
// 				"Terjadi kesalahan Sistem",
// 				false,
// 				err.Error(),
// 			),
// 		)
// 		return
// 	}

// 	c.JSON(
// 		http.StatusOK,
// 		utilities.ApiResponse(
// 			"Alamat berhasil diubah",
// 			true,
// 			alamat.GetAlamatByIdFormatter(*res),
// 		),
// 	)
// }

// func (h *AlamatHandler) HapusAlamat(c *gin.Context) {
// 	idRaw, _ := c.Params.Get("idAlamat")
// 	id, _ := strconv.ParseUint(idRaw, 10, 64)

// 	userIdInterface, _ := c.Get("userId")
// 	userId := uint64(userIdInterface.(float64))

// 	res, err := h.service.GetById(id)

// 	if err != nil {
// 		if strings.Contains(err.Error(), "record not found") {
// 			c.JSON(
// 				http.StatusNotFound,
// 				utilities.ApiResponse(
// 					"Alamat tidak ditemukan",
// 					false,
// 					nil,
// 				),
// 			)
// 		} else {
// 			c.JSON(
// 				http.StatusInternalServerError,
// 				utilities.ApiResponse(
// 					"Terjadi kesalahan Sistem",
// 					false,
// 					err.Error(),
// 				),
// 			)
// 		}
// 		return
// 	}

// 	if res.IDUser != userId {
// 		c.JSON(
// 			http.StatusForbidden,
// 			utilities.ApiResponse(
// 				"Anda tidak memiliki akses untuk mengubah alamat ini",
// 				false,
// 				nil,
// 			),
// 		)
// 		return
// 	}

// 	if err := h.service.Delete(res); err != nil {
// 		c.JSON(
// 			http.StatusInternalServerError,
// 			utilities.ApiResponse(
// 				"Terjadi kesalahan Sistem",
// 				false,
// 				err.Error(),
// 			),
// 		)
// 		return
// 	}

// 	c.JSON(
// 		http.StatusOK,
// 		utilities.ApiResponse(
// 			"Alamat berhasil dihapus",
// 			true,
// 			alamat.GetAlamatByIdFormatter(*res),
// 		),
// 	)
// }

// func (h *AlamatHandler) GetAlamatUser(c *gin.Context) {
// 	userIdInterface, _ := c.Get("userId")
// 	userId := uint64(userIdInterface.(float64))

// 	res, err := h.service.GetAllUserAddress(userId)

// 	if err != nil {
// 		c.JSON(
// 			http.StatusInternalServerError,
// 			utilities.ApiResponse(
// 				"Terjadi kesalahan Sistem",
// 				false,
// 				err.Error(),
// 			),
// 		)
// 		return
// 	}

// 	c.JSON(
// 		http.StatusOK,
// 		utilities.ApiResponse(
// 			"Alamat berhasil diambil",
// 			true,
// 			alamat.GetUserAlamatFormatter(res),
// 		),
// 	)
// }

// func (h *AlamatHandler) GetAlamatById(c *gin.Context) {
// 	idRaw, _ := c.Params.Get("idAlamat")
// 	id, _ := strconv.ParseUint(idRaw, 10, 64)

// 	userIdInterface, _ := c.Get("userId")
// 	userId := uint64(userIdInterface.(float64))

// 	res, err := h.service.GetById(id)

// 	if err != nil {
// 		if strings.Contains(err.Error(), "record not found") {
// 			c.JSON(
// 				http.StatusNotFound,
// 				utilities.ApiResponse(
// 					"Alamat tidak ditemukan",
// 					false,
// 					nil,
// 				),
// 			)
// 		} else {
// 			c.JSON(
// 				http.StatusInternalServerError,
// 				utilities.ApiResponse(
// 					"Terjadi kesalahan Sistem",
// 					false,
// 					err.Error(),
// 				),
// 			)
// 		}
// 		return
// 	}

// 	if res.IDUser != userId {
// 		c.JSON(
// 			http.StatusForbidden,
// 			utilities.ApiResponse(
// 				"Anda tidak memiliki akses untuk melihat alamat ini",
// 				false,
// 				nil,
// 			),
// 		)
// 		return
// 	}

// 	c.JSON(
// 		http.StatusOK,
// 		utilities.ApiResponse(
// 			"Alamat berhasil diambil",
// 			true,
// 			alamat.GetAlamatByIdFormatter(*res),
// 		),
// 	)
// }

// func (h *AlamatHandler) GetAlamatSeller(c *gin.Context) {
// 	idRaw, _ := c.Params.Get("idSeller")
// 	id, _ := strconv.ParseUint(idRaw, 10, 64)

// 	res, err := h.service.GetAllUserAddress(id)

// 	if err != nil {
// 		c.JSON(
// 			http.StatusInternalServerError,
// 			utilities.ApiResponse(
// 				"Terjadi kesalahan Sistem",
// 				false,
// 				err.Error(),
// 			),
// 		)
// 		return
// 	}

// 	alamatSeller := res[0]

// 	if alamatSeller.IsUser {
// 		c.JSON(
// 			http.StatusNotFound,
// 			utilities.ApiResponse(
// 				"Alamat tidak ditemukan",
// 				false,
// 				nil,
// 			),
// 		)
// 		return
// 	}

// 	c.JSON(
// 		http.StatusOK,
// 		utilities.ApiResponse(
// 			"Alamat Seller berhasil diambil",
// 			true,
// 			alamat.GetAlamatByIdFormatter(alamatSeller),
// 		),
// 	)
// }
