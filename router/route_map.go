package router

import (
	"proxy_srv/router/proxy_pool"
)

func registerRouteMap() {
	proxy_pool.Register(router)
}
