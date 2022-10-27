package global

import (
	"backend/core/console"
	"fmt"
)

func PrintCobraHelp() {
	usageStr := `欢迎使用 ` + console.Green(AppName+` `+Version) + ` 可以使用 ` + console.Red(`-h`) + ` 查看命令`
	fmt.Printf("%s\n", usageStr)
}
