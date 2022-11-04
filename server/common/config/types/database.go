package types

// DatabaseConfigure 数据库配置
type DatabaseConfigure struct {
	// Host            string
	// Port            int
	// User            string
	// Passwd          string
	// DBName          string
	Driver          string
	Source          string
	ConnMaxIdleTime int
	ConnMaxLifeTime int
	MaxIdleConns    int
	MaxOpenConns    int
}
