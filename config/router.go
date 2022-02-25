package config

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	Gin        *gin.Engine
	BaseRouter *gin.RouterGroup
}

func NewRouter(engine *gin.Engine) Router {
	return Router{Gin: engine, BaseRouter: engine.Group("/api/v1")}
}
