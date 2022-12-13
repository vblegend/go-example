package dto

import "server/common/models"

type MenuDTO struct {
	// 菜单ID
	ID int `json:"id"`
	// 菜单标识符
	Name string `json:"name"`
	// 菜单类型  0 路由 1 IFrame
	Type int `json:"type"`
	// 菜单标题
	Title string `json:"title"`
	// 菜单的图标
	Icon string `json:"icon"`
	// 菜单的url路径
	Path string `json:"path"`
	// 菜单排序
	Sort int `json:"-"`
	// 父级ID
	ParentID int `json:"-"`
	// 默认状态
	Opened bool `json:"opened"`
	// 子菜单
	Children []*MenuDTO `json:"children,omitempty"`
}

func (dto *MenuDTO) FromModel(model models.Menu) *MenuDTO {
	dto.ID = model.ID
	dto.Name = model.Name
	dto.Type = model.Type
	dto.Title = model.Title
	dto.Icon = model.Icon
	dto.Path = model.Path
	dto.Sort = model.Sort
	dto.Opened = model.Opened
	dto.ParentID = model.ParentID
	dto.Children = make([]*MenuDTO, 0)
	return dto
}

type SortBy []*MenuDTO

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i].Sort < a[j].Sort }
