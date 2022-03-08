package main

import (
	"clean-arch-2/alamat"
	"clean-arch-2/config"
	"clean-arch-2/daerah"
	"clean-arch-2/handlers"
	"clean-arch-2/kategori"
	"clean-arch-2/keranjang"
	"clean-arch-2/middlewares"
	"clean-arch-2/produk"
	"clean-arch-2/role"
	"clean-arch-2/user"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	//init database
	db, err := config.InitDB()
	if err != nil {
		panic(err)
	}

	//load .env server
	config.Init(".server")

	// engine := gin.Default()
	// engine.Static("/public/images/products", "./public/images/products")

	//init router
	router := config.NewRouter(gin.Default())
	router.Gin.Static("/public/images/users", "./public/images/users")
	router.Gin.Static("/public/images/products", "./public/images/products")

	//init middleware
	middleware := middlewares.Middleware{}
	// roleMiddleware := middlewares.RoleMiddleware{}

	//init role
	roleRepo := role.NewRoleRepository(db)
	roleService := role.NewRoleService(roleRepo)

	//init auth
	userRepo := user.NewUserRepository(db)
	userService := user.NewUserService(userRepo)
	authHandler := handlers.NewAuthHandler(router, userService, roleService)

	//init daerah
	daerahRepo := daerah.NewDaerahRepository(db)
	daerahService := daerah.NewDaerahService(daerahRepo)
	daerahHandler := handlers.NewDaerahHandler(
		router,
		daerahService,
		middleware,
	)

	//init alamat
	alamatRepo := alamat.NewAlamatRepository(db)
	alamatService := alamat.NewAlamatService(alamatRepo)
	alamatHandler := handlers.NewAlamatHandler(
		router,
		alamatService,
		roleService,
		middleware,
	)

	//init kategori
	kategoriRepo := kategori.NewKategoriRepository(db)
	kategoriService := kategori.NewKategoriService(kategoriRepo)
	kategoriHandler := handlers.NewKategoriHandler(
		router,
		kategoriService,
		middleware,
	)

	//init produk
	produkRepo := produk.NewProdukRepository(db)
	produkService := produk.NewProdukService(produkRepo)
	produkHandler := handlers.NewProdukHandler(
		router,
		produkService,
		kategoriService,
		middleware,
	)

	//init keranjang
	keranjangRepo := keranjang.NewKeranjangRepository(db)
	keranjangService := keranjang.NewKeranjangService(keranjangRepo)
	keranjangHandler := handlers.NewKeranjangHandler(
		router,
		keranjangService,
		userService,
		alamatService,
		produkService,
		middleware,
	)

	//setup router
	handlers.NewHandlers(
		authHandler,
		daerahHandler,
		alamatHandler,
		kategoriHandler,
		produkHandler,
		keranjangHandler,
	).Setup()

	router.Gin.Run(":" + os.Getenv("PORT"))
}
