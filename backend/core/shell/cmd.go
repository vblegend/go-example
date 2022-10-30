package shell

import (
	"backend/core/echo"
	"backend/core/log"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func ExeCommand(name string, args ...string) error {
	log.Info(fmt.Sprintf("[%s]=>%s %v", echo.Green("run"), echo.Yellow(name), echo.Cyan(strings.Join(args, " "))))
	pwd, _ := os.Getwd()
	exe := exec.Command(name, args...)
	exe.Dir = pwd
	var out bytes.Buffer
	var stderr bytes.Buffer
	exe.Stdout = &out
	exe.Stderr = &stderr
	err := exe.Run()
	if out.Len() > 0 {
		log.Infof("[%s]=>%s", echo.Magenta("out"), out.String())
	}
	if err != nil {
		log.Errorf("[%s]=>", echo.Red("err"), err.Error())
		log.Errorf("[%s]=>", echo.Red("err"), stderr.String())
		return err
	}
	return nil
}

func ProcessExist(trait ...interface{}) bool {
	pwd, _ := os.Getwd()
	sb := strings.Builder{}
	sb.WriteString("ps -ef | grep -v \"grep\"")
	for _, v := range trait {
		switch t := v.(type) {
		case string:
			{
				sb.WriteString(fmt.Sprintf(" | grep \"%v\"", t))
			}
		case int8, uint8, int16, uint16, int32, uint32, int64, uint64:
			{
				sb.WriteString(fmt.Sprintf(" | grep \" %d \"", t))
			}
		default:
			{
				sb.WriteString(fmt.Sprintf(" | grep %v ", v))
			}
		}
	}
	exe := exec.Command("/bin/bash", "-c", sb.String())
	exe.Dir = pwd
	var out bytes.Buffer
	exe.Stdout = &out
	exe.Run()
	return len(out.String()) > 0
}
