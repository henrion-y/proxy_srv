package set_proxy

import (
	"errors"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/proxy"
	"net/http"
	"net/url"
	"proxy_srv/proxy_host"
)

// 设置爬虫代理
func SetCollyProxy(c *colly.Collector, proxyType int) error {
	proxyHosts := proxy_host.GetProxyHosts(proxyType)
	if proxyHosts == nil {
		return errors.New("获取代理失败， 未启用代理")
	}
	rp, err := proxy.RoundRobinProxySwitcher(proxyHosts...)
	if err != nil {
		return err
	}
	c.SetProxyFunc(rp)
	return nil
}

// 获取代理
func getProxy(proxyType int) func(_ *http.Request) (*url.URL, error) {
	proxyHost := proxy_host.GetProxyHost(proxyType)
	if proxyHost == "" {
		return nil
	}
	return func(_ *http.Request) (*url.URL, error) {
		return url.Parse(proxyHost)
	}
}

// 设置http代理
func SetHttpProxy(c *http.Client, proxyType int) error {
	httpProxy := getProxy(proxyType)
	if httpProxy == nil {
		return errors.New("获取代理失败， 未启用代理")
	}
	c.Transport = &http.Transport{Proxy: httpProxy}
	return nil
}

// 设置http代理
func SetHttpTransportProxy(c *http.Transport, proxyType int) error {
	httpProxy := getProxy(proxyType)
	if httpProxy == nil {
		return errors.New("获取代理失败， 未启用代理")
	}
	c = &http.Transport{Proxy: httpProxy}
	return nil
}
