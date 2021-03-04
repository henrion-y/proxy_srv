package proxy_host

import (
	"testing"
	"time"
)

func TestAddProxyHost(t *testing.T) {
	AddProxyHost("http://192.168.9.1:12345", FOREIGN_PROXY)
	t.Log(GetProxyHosts(FOREIGN_PROXY))
	AddProxyHost("http://192.168.9.1:6785", FOREIGN_PROXY)
	t.Log(GetProxyHosts(FOREIGN_PROXY))
	AddProxyHost("http://127.0.0.1:1090", FOREIGN_PROXY)
	t.Log(GetProxyHosts(FOREIGN_PROXY))
}

func TestAddProxyHosts(t *testing.T) {
	hosts := []string{
		"http://192.168.9.1:12345",
		"http://192.168.9.1:1561",
		"http://192.168.9.1:53453",
		"http://192.168.9.1:5686",
		"http://192.168.9.1:890",
		"http://127.0.0.1:1090",
	}
	go AddProxyHosts(hosts, FOREIGN_PROXY)
	time.Sleep(2*time.Second)
	hosts2 := []string{
		"http://192.168.9.1:545",
		"http://192.168.9.1:181",
		"http://192.168.9.1:513",
		"http://192.168.9.1:5086",
		"http://192.168.9.1:8860",
		"http://127.0.0.1:1090",
	}
	AddProxyHosts(hosts2, FOREIGN_PROXY)
	t.Log(GetProxyHosts(FOREIGN_PROXY))
}

func TestDelProxyHost(t *testing.T) {
	hosts := []string{
		"http://192.168.9.1:12345",
		"http://192.168.9.1:1561",
		"http://192.168.9.1:53453",
		"http://192.168.9.1:5686",
		"http://192.168.9.1:890",
	}
	AddProxyHosts(hosts, FOREIGN_PROXY)
	t.Log(GetProxyHosts(FOREIGN_PROXY))
	DelProxyHost("http://192.168.9.1:12345", FOREIGN_PROXY)
	t.Log(GetProxyHosts(FOREIGN_PROXY))
}

func TestGetProxyHost(t *testing.T) {
	hosts := []string{
		"http://192.168.9.1:12345",
		"http://192.168.9.1:1561",
		"http://192.168.9.1:53453",
		"http://192.168.9.1:5686",
		"http://192.168.9.1:890",
	}
	AddProxyHosts(hosts, FOREIGN_PROXY)
	t.Log(GetProxyHost(FOREIGN_PROXY))
	DelProxyHost("http://192.168.9.1:12345", FOREIGN_PROXY)
	t.Log(GetProxyHost(FOREIGN_PROXY))
}

func TestClearBadProxyHosts(t *testing.T) {
	hosts := []string{
		"http://192.168.9.1:12345",
		"http://192.168.9.1:1561",
		"http://192.168.9.1:53453",
		"http://192.168.9.1:5686",
		"http://192.168.9.1:890",
		"http://192.168.9.1:890",
		"http://127.0.0.1:1090",
	}
	AddProxyHosts(hosts, FOREIGN_PROXY)
	t.Log("清理前：", GetProxyHosts(FOREIGN_PROXY))
	go ClearBadProxyHosts()
	ClearBadProxyHosts()
	t.Log("清理后：", GetProxyHosts(FOREIGN_PROXY))
}
