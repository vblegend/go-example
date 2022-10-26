package pkg

import (
	"backend/common/global"
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var tempFile = fmt.Sprintf("/tmp/%s.pid", global.AppFileName)

func RemovePidFile() {
	os.Remove(tempFile)
}

func GetProcessByCommand(cmd string) []string {
	pwd, _ := os.Getwd()
	exe := exec.Command("/bin/bash", "-c", fmt.Sprintf("ps -ef | grep -v \"grep\" | grep \"%s\" | awk '{print $2}'", cmd))
	exe.Dir = pwd
	var out bytes.Buffer
	exe.Stdout = &out
	exe.Run()
	if len(out.String()) > 0 {
		pids := strings.Split(out.String(), "\n")
		return pids[:len(pids)-1]
	}
	return []string{}
}

func SelfProcessExist(pid int) bool {
	pwd, _ := os.Getwd()
	exe := exec.Command("/bin/bash", "-c", fmt.Sprintf("ps -ef | grep -v \"grep\" | grep \"./%s server\" | grep \" %d \"", global.AppFileName, pid))
	exe.Dir = pwd
	var out bytes.Buffer
	exe.Stdout = &out
	exe.Run()
	fmt.Print(out.String())
	return len(out.String()) > 0
}

func RunOfOnec() error {
	if FileExist(tempFile) {
		data, err := os.ReadFile(tempFile)
		if err != nil {
			os.Remove(tempFile)
		}
		pid, err := strconv.Atoi(string(data))
		if err != nil {
			os.Remove(tempFile)
		}
		if SelfProcessExist(pid) {
			return errors.New("Process is only allowed to run once")
		}
	}
	os.WriteFile(tempFile, []byte(fmt.Sprintf("%d", os.Getpid())), os.ModePerm)
	return nil
}

func IsRuning(pidOut *int) bool {
	if FileExist(tempFile) {
		data, err := os.ReadFile(tempFile)
		if err != nil {
			os.Remove(tempFile)
		}
		pid, err := strconv.Atoi(string(data))
		if err != nil {
			os.Remove(tempFile)
		}
		if SelfProcessExist(pid) {
			*pidOut = pid
			return true
		}
	}
	return false
}
