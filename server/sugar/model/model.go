package model

import "gorm.io/gorm"

// IModel 数据库表模型对象接口,可选择性组合以下任意功能性接口
//
// IModelQueryAfter,IModelInsertBefore,IModelInsertAfter,IModelUpdateBefore,IModelUpdateAfter,IModelDeleteBefore,IModelDeleteAfter
type IModel interface {
	// TableName 模型对应表明
	TableName() string
}

// IModelQueryAfter 【IModel组合 QueryAfter Hook接口】
type IModelAfterFind interface {
	// 数据从数据表中查询出来之后触发
	AfterFind(tx *gorm.DB) (err error)
}

// IModelInsertBefore 【IModel组合 InsertBefore Hook接口】
type IModelBeforeCreate interface {
	// OnInsertBefore 数据插入数据表之前触发
	BeforeCreate(tx *gorm.DB) (err error)
}

// IModelInsertAfter 【IModel组合 InsertAfter Hook接口】
type IModelAfterCreate interface {
	// OnInsertAfter 数据插入数据表之后触发
	AfterCreate(tx *gorm.DB) (err error)
}

// IModelUpdateBefore 【IModel组合 UpdateBefore Hook接口】
type IModelBeforeUpdate interface {
	// OnUpdateBefore 数据更新至数据表之前触发
	BeforeUpdate(tx *gorm.DB) (err error)
}

// IModelUpdateAfter 【IModel组合 UpdateAfter Hook接口】
type IModelAfterUpdate interface {
	// OnUpdateAfter 数据更新至数据表之后触发
	AfterUpdate(tx *gorm.DB) (err error)
}

// IModelDeleteBefore 【IModel组合 DeleteBefore Hook接口】
type IModelBeforeDelete interface {
	// OnDeleteBefore 数据从数据表删除之前触发
	BeforeDelete(tx *gorm.DB) (err error)
}

// IModelDeleteAfter 【IModel组合 DeleteAfter Hook接口】
type IModelAfterDelete interface {
	// OnDeleteAfter 数据从数据表删除之后触发
	AfterDelete(tx *gorm.DB) (err error)
}
