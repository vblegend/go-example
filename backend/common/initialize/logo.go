package initialize

import (
	"backend/common/assembly"
	"backend/common/config"
	"backend/core/echo"
	"backend/core/log"
	"fmt"
	"strings"
)

func PrintLogo() {
	log.Print(echo.Yellow(strings.Join(assembly.LogoContent, "\n")) + fmt.Sprintf(" %s %s (%s)\n", echo.Green(config.Application.Mode), echo.Red("V"+assembly.Version), assembly.BuildTime))

}
