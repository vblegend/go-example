package model

// IModel 数据库表模型对象接口,可选择性组合以下任意功能性接口
//
// IModelQueryAfter,IModelInsertBefore,IModelInsertAfter,IModelUpdateBefore,IModelUpdateAfter,IModelDeleteBefore,IModelDeleteAfter
type IModel interface {
	// TableName 模型对应表明
	TableName() string
}

// IModelQueryAfter 【IModel组合 QueryAfter Hook接口】
type IModelQueryAfter interface {
	// 数据从数据表中查询出来之后触发
	OnQueryAfter()
}

// IModelInsertBefore 【IModel组合 InsertBefore Hook接口】
type IModelInsertBefore interface {
	// OnInsertBefore 数据插入数据表之前触发
	OnInsertBefore() error
}

// IModelInsertAfter 【IModel组合 InsertAfter Hook接口】
type IModelInsertAfter interface {
	// OnInsertAfter 数据插入数据表之后触发
	OnInsertAfter()
}

// IModelUpdateBefore 【IModel组合 UpdateBefore Hook接口】
type IModelUpdateBefore interface {
	// OnUpdateBefore 数据更新至数据表之前触发
	OnUpdateBefore() error
}

// IModelUpdateAfter 【IModel组合 UpdateAfter Hook接口】
type IModelUpdateAfter interface {
	// OnUpdateAfter 数据更新至数据表之后触发
	OnUpdateAfter()
}

// IModelDeleteBefore 【IModel组合 DeleteBefore Hook接口】
type IModelDeleteBefore interface {
	// OnDeleteBefore 数据从数据表删除之前触发
	OnDeleteBefore() error
}

// IModelDeleteAfter 【IModel组合 DeleteAfter Hook接口】
type IModelDeleteAfter interface {
	// OnDeleteAfter 数据从数据表删除之后触发
	OnDeleteAfter()
}
