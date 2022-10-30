package initialize

import (
	"backend/common/config"
	"backend/core/env"
	"os"
)

func InitRunMode() {
	// 运行模式， 开发环境 或 生产环境
	// 如果命令行指定了环境变量 将覆盖配置文件中的选项
	env.SetMode(env.ParseMode(config.Application.Mode))
	RUN_MODE := os.Getenv("RUN_MODE")
	if len(RUN_MODE) > 0 {
		env.SetMode(env.ParseMode(RUN_MODE))
	}
}
