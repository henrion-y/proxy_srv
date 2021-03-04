package form

type GetHostReq struct {
	HostType  int `json:"host_type" form:"host_type"`
	ProxyType int `json:"proxy_type" form:"proxy_type"`
}
