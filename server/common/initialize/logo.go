package initialize

import (
	"fmt"
	"server/common/assembly"
	"server/common/config"
	"server/sugar/echo"
	"server/sugar/log"
	"strings"
)

func PrintLogo() {
	log.Print(echo.Yellow(strings.Join(assembly.LogoContent, "\n")) + fmt.Sprintf(" %s %s (%s)", echo.Green(config.Application.Mode), echo.Red("V"+assembly.Version), assembly.BuildTime))

}
