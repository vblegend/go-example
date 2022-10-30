package types

type RedisConfigure struct {
	Host     string `yaml:"host" form:"host" json:"host"`
	Port     int32  `yaml:"port" form:"port" json:"port"`
	Password string `yaml:"password" form:"password" json:"password"`
	DB       int    `yaml:"db" form:"db" json:"db"`
}
