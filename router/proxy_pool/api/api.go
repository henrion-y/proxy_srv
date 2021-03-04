package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"proxy_srv/lib/helper"
	"proxy_srv/proxy_host"
	"proxy_srv/router/proxy_pool/form"
)

func Mapping(prefix string, app *gin.Engine) {
	router := app.Group(prefix)
	router.GET("/host", getHost)
}

func getHost(ctx *gin.Context) {
	var param form.GetHostReq
	err := ctx.ShouldBindWith(&param, binding.Query)
	if err != nil {
		ctx.JSON(helper.Fail(err.Error()))
		return
	}
	if param.ProxyType != proxy_host.FOREIGN_PROXY && param.ProxyType != proxy_host.DOMESTIC_PROXY {
		ctx.Status(400)
		return
	}

	if param.HostType != 0 && param.HostType != 1 {
		ctx.Status(400)
		return
	}

	var result form.GetHostResp
	if param.HostType == 0 {
		host := proxy_host.GetProxyHost(param.ProxyType)
		result.Hosts = append(result.Hosts, host)
	} else {
		result.Hosts = proxy_host.GetProxyHosts(param.ProxyType)
	}

	ctx.JSON(helper.SuccessWithData(result))
	return
}
