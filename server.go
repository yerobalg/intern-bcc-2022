package main

import (
	"github.com/gin-gonic/gin"
	// c "clean-arch-2/controllers"
	"clean-arch-2/app/handlers"
	"clean-arch-2/config"
	"clean-arch-2/app/services"
	"clean-arch-2/app/repositories"
	//"os"
	// "clean-arch-2/entities"
	// model "clean-arch-2/models"
	// "fmt"
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
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	authHandler := handlers.NewAuthHandler(router, userService)

	//setup router
	handlers.NewHandlers(
		authHandler,
	).Setup()
	
	router.Gin.Run()
}
