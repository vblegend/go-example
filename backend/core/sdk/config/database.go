package config

type Database struct {
	Driver          string
	Source          string
	ConnMaxIdleTime int
	ConnMaxLifeTime int
	MaxIdleConns    int
	MaxOpenConns    int
}

var (
	DatabaseConfig = map[string]*Database{}
)
