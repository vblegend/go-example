package config

type Mysql struct {
	Host     string `yaml:"host" form:"host" json:"host"`
	Port     int    `yaml:"port" form:"port" json:"port"`
	User     string `yaml:"user" form:"user" json:"user"`
	Password string `yaml:"password" form:"password" json:"password"`
}

var MysqlConfig = new(Mysql)
