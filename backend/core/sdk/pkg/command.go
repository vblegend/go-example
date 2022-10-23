package pkg

import (
	"backend/core/logger"
	"backend/core/sdk/console"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func ExeCommand(name string, args ...string) error {
	logger.Info(fmt.Sprintf("[%s]=>%s %v", console.Green("run"), console.Yellow(name), console.Cyan(strings.Join(args, " "))))
	pwd, _ := os.Getwd()
	exe := exec.Command(name, args...)
	exe.Dir = pwd
	var out bytes.Buffer
	var stderr bytes.Buffer
	exe.Stdout = &out
	exe.Stderr = &stderr
	err := exe.Run()
	if out.Len() > 0 {
		logger.Infof("[%s]=>%s", console.Magenta("out"), out.String())
	}
	if err != nil {
		logger.Errorf("[%s]=>", console.Red("err"), err.Error())
		logger.Errorf("[%s]=>", console.Red("err"), stderr.String())
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
