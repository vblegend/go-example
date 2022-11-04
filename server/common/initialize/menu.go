package initialize

import (
	"server/common/config"
	"server/sugar/env"
	"server/sugar/state"
)

func InitDevelopmentMenu() {
	db := state.Default.GetDB(config.DefaultDB)
	if db != nil {
		visible := 0
		if env.ModeIs(env.Production) {
			visible = 1
		} // 开发模式下显示菜单配置页面
		db.Exec("UPDATE sys_menu SET visible = ? WHERE menu_id = 51", visible)
	}
}
