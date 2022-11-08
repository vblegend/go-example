package service

import (
	"server/app/admin/dto"
	"server/common/models"
	"server/sugar/service"
	"sort"
)

type MenuService struct {
	service.Service
}

func (s *MenuService) GetMenuTree() []*dto.MenuDTO {
	menus := models.Menus{}
	err := s.Orm.Table(models.Menu{}.TableName()).Where("visible = ?", 1).Find(&menus).Error
	if err != nil {
		// api.Error(http.StatusInternalServerError, err)
		// return nil
	}
	dtoMaps := make(map[int]*dto.MenuDTO)
	// roots := menus.Roots()
	for _, menu := range menus {
		menuDTO := new(dto.MenuDTO).FromModel(menu)
		dtoMaps[menuDTO.ID] = menuDTO
	}
	for _, menu := range dtoMaps {
		if menu.ParentID > 0 {
			if parentMenu, ok := dtoMaps[menu.ParentID]; ok {
				parentMenu.Children = append(parentMenu.Children, menu)
			}
		}
	}
	result := make([]*dto.MenuDTO, 0)
	for _, menu := range dtoMaps {
		if menu.ParentID == 0 {
			if len(menu.Children) > 0 {
				sort.Sort(dto.SortBy(menu.Children))
				// menu.Children
			}
			result = append(result, menu)
		}
	}
	updateMenuPath("", result)
	sort.Sort(dto.SortBy(result))
	return result
}

func updateMenuPath(basePath string, menus []*dto.MenuDTO) {
	for _, menu := range menus {
		menu.Path = basePath + menu.Path
		updateMenuPath(menu.Path, menu.Children)
	}
}
