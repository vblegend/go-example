package assembly

import (
	"fmt"
	"server/sugar/echo"
)

func PrintCobraHelp() {
	usageStr := `欢迎使用 ` + echo.Green(AppName+` `+Version) + ` 可以使用 ` + echo.Red(`-h`) + ` 查看命令`
	fmt.Printf("%s\n", usageStr)
}
