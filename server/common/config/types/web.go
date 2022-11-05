package types

// WebConfigure web服务配置
type WebConfigure struct {
	Host string
	Port int64
	// Https 是否启用Https
	Https    bool
	Root     string
	CertFile string
	KeyFile  string
	Domain   string
	// WhiteList []string
	// BlackList []string
}

// GetHTTPProtocol 获取当前http协议
func (a WebConfigure) GetHTTPProtocol() string {
	if a.Https {
		return "Https"
	} else {
		return "Http"
	}
}
