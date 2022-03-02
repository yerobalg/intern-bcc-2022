package handlers

import (
	"clean-arch-2/config"
	"clean-arch-2/alamat"
	"clean-arch-2/middlewares"
	// "clean-arch-2/utilities"
	// "fmt"
	// "github.com/gin-gonic/gin"
	// "net/http"
	// "strings"
)

type AlamatHandler struct {
	handler    config.Router
	service    alamat.AlamatService
	middleware middlewares.Middleware
}

func (h AlamatHandler) Setup() {
	api := h.handler.BaseRouter.Use(h.middleware.AuthMiddleware())
	{
		api.POST("/alamat",)
		api.GET("/alamat/:idAlamat",)
		api.PATCH("/alamat/:idAlamat", )
		api.DELETE("/alamat/:idAlamat", )
	}
}


func NewAlamatHandler(
	handler config.Router,
	service alamat.AlamatService,
	middleware middlewares.Middleware,
) AlamatHandler {
	return AlamatHandler{
		handler:    handler,
		service:    service,
		middleware: middleware,
	}
}