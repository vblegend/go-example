package models

type Menu struct {
	// 菜单ID
	ID int `json:"id" gorm:"primaryKey;autoIncrement"`
	// 菜单标识符
	Name string `json:"name" gorm:"size:32;"`
	// 菜单类型  0 路由 1 IFrame
	Type int `json:"type" gorm:"size:2;DEFAULT:0;"`
	// 菜单标题
	Title string `json:"title" gorm:"size:64;"`
	// 菜单的图标
	Icon string `json:"icon" gorm:"size:32;"`
	// 菜单的url路径
	Path string `json:"path" gorm:"size:128;"`
	// 上级菜单ID
	ParentID int `json:"parentId" gorm:"size:11;"`
	// 菜单排序
	Sort int `json:"sort" gorm:"size:4;"`
	// 菜单是否可见
	Visible string `json:"visible" gorm:"size:1;"`
	//
	ModelTime
}

func (Menu) TableName() string {
	return "menu"
}

type Menus []Menu

func (ms Menus) Roots() Menus {
	roots := Menus{}
	for _, menu := range ms {
		if menu.ParentID == 0 {
			roots = append(roots, menu)
		}
	}
	return roots
}
