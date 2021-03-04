package router

import (
	"github.com/gin-gonic/gin"
	"proxy_srv/lib/middleware/cors"
)

var router *gin.Engine

func Init() *gin.Engine {
	gin.SetMode("debug")
	router = gin.New()
	router.Use(gin.Recovery(), gin.Logger(), cors.CORS())
	registerRouteMap()
	return router
}
