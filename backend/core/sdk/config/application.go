package config

type Application struct {
	Host string
	Port int64
	Name string
	Mode string
}

var ApplicationConfig = new(Application)
