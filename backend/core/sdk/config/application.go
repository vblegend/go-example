package config

type Application struct {
	Host     string
	Port     int64
	Name     string
	Mode     string
	Https    bool
	CertFile string
	KeyFile  string
	Domain   string
}

func (a Application) GetHttpProtocol() string {
	if a.Https {
		return "Https"
	} else {
		return "Http"
	}
}

var ApplicationConfig = new(Application)
