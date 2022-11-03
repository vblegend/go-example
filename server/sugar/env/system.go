package env

import (
	"io/ioutil"
	"os"
	"runtime"
	"strings"
)

const (
	Linux_Ubuntu  = "Ubuntu"
	Linux_CentOS  = "CentOS"
	Linux_Debian  = "Debian"
	Linux_Unknown = "Unknown"
	Windows       = "Windows"
	osReleaseFile = "/etc/os-release"
)

var system = Windows

func init() {
	if runtime.GOOS == "linux" {
		system = Linux_Unknown
		if info, err := os.Stat(osReleaseFile); err == nil {
			if !info.IsDir() {
				if data, err := ioutil.ReadFile(osReleaseFile); err == nil {
					content := strings.ToLower(string(data))
					if strings.Index(content, "ubuntu") > 0 {
						system = Linux_Ubuntu
					} else if strings.Index(content, "centos") > 0 {
						system = Linux_CentOS
					} else if strings.Index(content, "debian") > 0 {
						system = Linux_Debian
					}
				}
			}
		}
	} else if runtime.GOOS == "windows" {
		system = Windows
	}
}

func System() string {
	return system
}
