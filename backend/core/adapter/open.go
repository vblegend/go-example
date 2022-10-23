package gormadapter

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var opens = map[string]func(string) gorm.Dialector{
	"mysql":   mysql.Open,
	"sqlite3": sqlite.Open,
}
