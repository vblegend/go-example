package types

type ApplicationConfigure struct {
	Host     string
	Port     int64
	Name     string
	Mode     string
	Https    bool
	CertFile string
	KeyFile  string
	Domain   string
}

func (a ApplicationConfigure) GetHttpProtocol() string {
	if a.Https {
		return "Https"
	} else {
		return "Http"
	}
}
