package models

import "time"

// Migration 数据迁移版本表
type Migration struct {
	Version   string    `gorm:"primaryKey"`
	ApplyTime time.Time `gorm:"autoCreateTime"`
}

// TableName 数据表名称
func (Migration) TableName() string {
	return "sys_migration"
}

// Migrations  Migration的切片集合
type Migrations []Migration

// Map 把切片转换为 Version Map
func (ms Migrations) Map() map[string]*Migration {
	m := make(map[string]*Migration)
	for _, i := range ms {
		m[i.Version] = &i
	}
	return m
}
