package proxy_srv

import (
	"fmt"
	"proxy_srv/pool_services/colly_kuaidaili_proxy"
	"proxy_srv/proxy_host"
	"proxy_srv/router"
	"runtime/debug"
	"time"
)

type ProxySrv struct {
	// 这里放各个站点的开关， 默认开启
	IgnoreKuaidaili bool

	// 这里是爬取各站点的频率
	CrawlFrequency time.Duration

	// 这里是检查失效的代理的频率
	ClearHostFrequency time.Duration

	Debug  bool
	RunWeb bool
}

type cronFunc func()

func start(f cronFunc) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("任务异常 ： ", r)
			fmt.Println(string(debug.Stack()))
		}
	}()
	f()
}

func cron(d time.Duration, f cronFunc) {
	for {
		start(f)
		time.Sleep(d)
	}
}

func (s *ProxySrv) crawl() {
	if !s.IgnoreKuaidaili {
		kuaidailiSrv := colly_kuaidaili_proxy.CollySrv{
			UseProxy: true,
		}
		kuaidailiSrv.Run()
	}
}

func (s *ProxySrv) Run() {
	go cron(s.CrawlFrequency, s.crawl)
	proxy_host.Debug = s.Debug
	go cron(s.ClearHostFrequency, proxy_host.ClearBadProxyHosts)
	if s.RunWeb {
		web := router.Init()
		err := web.Run(":8081")
		if err != nil {
			panic(err)
		}
	}
}
