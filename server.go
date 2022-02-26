package main

import (
	"github.com/gin-gonic/gin"
	"clean-arch-2/handlers"
	"clean-arch-2/user"
	"clean-arch-2/daerah"
	"clean-arch-2/config"
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

	//init auth
	userRepo := user.NewUserRepository(db)
	userService := user.NewUserService(userRepo)
	authHandler := handlers.NewAuthHandler(router, userService)

	//init daerah
	daerahRepo := daerah.NewDaerahRepository(db)
	daerahService := daerah.NewDaerahService(daerahRepo)
	daerahHandler := handlers.NewDaerahHandler(router, daerahService)

	//setup router
	handlers.NewHandlers(
		authHandler,
		daerahHandler,
	).Setup()

	router.Gin.Run()
}
