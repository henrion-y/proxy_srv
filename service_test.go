package proxy_srv

import (
	"testing"
	"time"
)

func TestProxySrv_Run(t *testing.T) {
	proxySrv := ProxySrv{
		CrawlFrequency:     3 * time.Minute,
		ClearHostFrequency: 2 * time.Minute,
		RunWeb:             true,
	}
	proxySrv.Run()
}
