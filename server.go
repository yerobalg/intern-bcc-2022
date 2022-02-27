package main

import (
	"clean-arch-2/config"
	"clean-arch-2/daerah"
	"clean-arch-2/handlers"
	"clean-arch-2/middlewares"
	"clean-arch-2/user"

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
	authMiddleware := middlewares.AuthMiddleware{}
	// roleMiddleware := middlewares.RoleMiddleware{}

	//init auth
	userRepo := user.NewUserRepository(db)
	userService := user.NewUserService(userRepo)
	authHandler := handlers.NewAuthHandler(router, userService)

	//init daerah
	daerahRepo := daerah.NewDaerahRepository(db)
	daerahService := daerah.NewDaerahService(daerahRepo)
	daerahHandler := handlers.NewDaerahHandler(
		router,
		daerahService,
		authMiddleware,
	)

	//setup router
	handlers.NewHandlers(
		authHandler,
		daerahHandler,
	).Setup()

	router.Gin.Run()
}
