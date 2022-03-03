package main

import (
	"clean-arch-2/config"
	"clean-arch-2/handlers"
	"clean-arch-2/middlewares"
	"clean-arch-2/role"
	"clean-arch-2/user"
	"clean-arch-2/daerah"
	"clean-arch-2/alamat"
	"clean-arch-2/kategori"
	"github.com/gin-gonic/gin"
)

func main() {
	//init database
	db, err := config.InitDB()
	if err != nil {
		panic(err)
	}

	//load .env server
	config.Init(".server")

	//init router
	router := config.NewRouter(gin.Default())
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

	//setup router
	handlers.NewHandlers(
		authHandler,
		daerahHandler,
		alamatHandler,
		kategoriHandler,
	).Setup()

	router.Gin.Run()
}
