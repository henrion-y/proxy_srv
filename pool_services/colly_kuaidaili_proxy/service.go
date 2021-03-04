package colly_kuaidaili_proxy

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
	"log"
	"proxy_srv/proxy_host"
	"proxy_srv/set_proxy"
	"strings"
)

var kuaidailiPages = []string{
	"https://www.kuaidaili.com/free/",
}

type CollySrv struct {
	CrawledUrls []string
	UseProxy    bool
	Debug       bool
}

func (s *CollySrv) getTableHosts(sourceUrl string) {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36"))

	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	// 使用代理进行爬取
	if s.UseProxy {
		err := set_proxy.SetCollyProxy(c, proxy_host.FOREIGN_PROXY)
		if err != nil {
			if s.Debug {
				fmt.Println(err)
			}
		}
	}

	c.OnHTML("tbody", func(element *colly.HTMLElement) {
		var hosts []string
		element.ForEach("tr", func(_ int, elem *colly.HTMLElement) {
			ip := elem.ChildText("td[data-title=IP]")
			port := elem.ChildText("td[data-title=PORT]")
			scheme := elem.ChildText("td[data-title=类型]")
			host := fmt.Sprintf("%s://%s:%s", strings.ToLower(scheme), ip, port)
			hosts = append(hosts, host)
		})
		proxy_host.AddProxyHosts(hosts, proxy_host.FOREIGN_PROXY)
	})

	err := c.Visit(sourceUrl)
	if err != nil {
		log.Println(err)
	}
}

func (s *CollySrv) Run() {
	if s.CrawledUrls == nil {
		s.CrawledUrls = kuaidailiPages
	}
	for _, v := range s.CrawledUrls {
		s.getTableHosts(v)
	}
}
