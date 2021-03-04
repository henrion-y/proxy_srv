package colly_kuaidaili_proxy

import (
	"fmt"
	"proxy_srv/proxy_host"
	"testing"
)

func TestGetPageList(t *testing.T) {
	kuaidailiSrv := CollySrv{
		UseProxy: true,
	}
	kuaidailiSrv.Run()
	fmt.Println("抓取到的代理为： ", proxy_host.GetProxyHosts(proxy_host.FOREIGN_PROXY))
}
