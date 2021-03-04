package proxy_host

import (
	"fmt"
	"net/http"
	"net/url"
	"sync"
)

const (
	baiduDesc  = "https://www.baidu.com/"
	googleDesc = "https://www.google.com/"

	FOREIGN_PROXY  = 0
	DOMESTIC_PROXY = 1
)

var Debug bool

type proxyHosts struct {
	data map[string]struct{}
	lock sync.RWMutex
}

var domesticProxyHosts = &proxyHosts{data: make(map[string]struct{})}
var foreignProxyHosts = &proxyHosts{data: make(map[string]struct{})}

func CheckProxyHost(host string, proxyType int) bool {
	destAddr := googleDesc
	if proxyType == DOMESTIC_PROXY {
		destAddr = baiduDesc
	}
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse(host)
	}
	transport := &http.Transport{Proxy: proxy}
	client := &http.Client{Transport: transport}
	_, err := client.Get(destAddr)
	if err != nil {
		return false
	}
	return true
}

func addToHosts(host string, proxyType int, async bool, addChWg chan<- struct{}) bool {
	defer func() {
		if async {
			addChWg <- struct{}{}
		}
	}()

	if ok := CheckProxyHost(host, proxyType); !ok {
		return false
	}

	proxyHosts := foreignProxyHosts
	if proxyType == DOMESTIC_PROXY {
		proxyHosts = domesticProxyHosts
	}

	proxyHosts.lock.Lock()
	defer proxyHosts.lock.Unlock()
	proxyHosts.data[host] = struct{}{}
	return true
}

func AddProxyHost(host string, proxyType int) bool {
	return addToHosts(host, proxyType, false, nil)
}

func AddProxyHosts(hosts []string, proxyType int) {
	addChWg := make(chan struct{}, len(hosts))
	for _, host := range hosts {
		go addToHosts(host, proxyType, true, addChWg)
	}
	for i := 0; i < len(hosts); i++ {
		<-addChWg
	}
}

func GetProxyHosts(proxyType int) []string {
	var hosts []string
	proxyHosts := foreignProxyHosts
	if proxyType == DOMESTIC_PROXY {
		proxyHosts = domesticProxyHosts
	}

	proxyHosts.lock.RLock()
	defer proxyHosts.lock.RUnlock()
	for host := range proxyHosts.data {
		hosts = append(hosts, host)
	}
	return hosts
}

func GetProxyHost(proxyType int) string {
	proxyHosts := foreignProxyHosts
	if proxyType == DOMESTIC_PROXY {
		proxyHosts = domesticProxyHosts
	}

	proxyHosts.lock.RLock()
	defer proxyHosts.lock.RUnlock()
	for host := range proxyHosts.data {
		return host
	}
	return ""
}

func DelProxyHost(host string, proxyType int) {
	proxyHosts := foreignProxyHosts
	if proxyType == DOMESTIC_PROXY {
		proxyHosts = domesticProxyHosts
	}

	proxyHosts.lock.Lock()
	defer proxyHosts.lock.Unlock()
	delete(proxyHosts.data, host)
}

func clear(host string, proxyType int, clearChWg chan<- struct{}) bool {
	defer func() {
		clearChWg <- struct{}{}
	}()

	if ok := CheckProxyHost(host, proxyType); ok {
		return false
	}

	proxyHosts := foreignProxyHosts
	if proxyType == DOMESTIC_PROXY {
		proxyHosts = domesticProxyHosts
	}

	proxyHosts.lock.Lock()
	defer proxyHosts.lock.Unlock()
	delete(proxyHosts.data, host)
	return true
}

// 清理代理池中失效的代理
func ClearBadProxyHosts() {
	defer func() {
		if Debug {
			fmt.Println("清理完毕")
			fmt.Println("现有国内ip池 ： ", GetProxyHosts(DOMESTIC_PROXY))
			fmt.Println("现有国外ip池 ： ", GetProxyHosts(FOREIGN_PROXY))
		}
	}()

	chLen := len(domesticProxyHosts.data) + len(foreignProxyHosts.data)
	clearChWg := make(chan struct{}, chLen)

	for host := range domesticProxyHosts.data {
		go clear(host, DOMESTIC_PROXY, clearChWg)
	}
	for host := range foreignProxyHosts.data {
		go clear(host, FOREIGN_PROXY, clearChWg)
	}
	for i := 0; i < chLen; i++ {
		<-clearChWg
	}
}
