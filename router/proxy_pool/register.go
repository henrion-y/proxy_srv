package proxy_pool

import (
	"github.com/gin-gonic/gin"
	"proxy_srv/router/proxy_pool/api"
)

const (
	defaultRoutePrefix = "/proxy_pool"
)

func Register(app *gin.Engine) {
	api.Mapping(defaultRoutePrefix, app)
}
