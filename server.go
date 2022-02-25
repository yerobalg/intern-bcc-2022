package main

import (
	// "github.com/gin-gonic/gin"
	// c "clean-arch-2/controllers"
	// "clean-arch-2/app/handlers"
	"clean-arch-2/config"
	// "clean-arch-2/app/services"
	// "clean-arch-2/app/repositories"
	//"os"
	// "clean-arch-2/entities"
	// model "clean-arch-2/models"
	// "fmt"
)

func main() {
	_, err := config.InitDB()
	if err != nil {
		panic(err)
	}

	// config.Init(".server")

	// router := config.NewRouter(gin.Default())

	// postRepo := repositories.NewPostRepository(db)
	// postService := services.NewPostService(postRepo)
	// postHandler := handlers.NewPostHandler(router, postService)

	// handlers.NewHandlers(
	// 	postHandler,
	// ).Setup()
	
	// router.Gin.Run()

	// router.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// router.GET("/post", c.GetPosts)
	// router.POST("/post", c.AddNewPost)
	// conf.InitDB()
	// db, err := model.InitDB(
	// 	conf.GetEnv("DB_USERNAME"),
	// 	conf.GetEnv("DB_DBNAME"),
	// 	conf.GetEnv("DB_PASSWORD"),
	// 	conf.GetEnv("DB_HOST"),
	// 	conf.GetEnv("DB_PORT"),
	// )

	// if err != nil {

	// }

}
