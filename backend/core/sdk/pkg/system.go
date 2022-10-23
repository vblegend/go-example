package pkg

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

var current_system = Windows

func init() {
	if runtime.GOOS == "linux" {
		if info, err := os.Stat(osReleaseFile); err == nil {
			if !info.IsDir() {
				if data, err := ioutil.ReadFile(osReleaseFile); err == nil {
					content := strings.ToLower(string(data))
					if strings.Index(content, "ubuntu") > 0 {
						current_system = Linux_Ubuntu
					} else if strings.Index(content, "centos") > 0 {
						current_system = Linux_CentOS
					} else if strings.Index(content, "debian") > 0 {
						current_system = Linux_Debian
					} else {
						current_system = Linux_Unknown
					}
				}
			}
		}
	} else if runtime.GOOS == "windows" {
		current_system = Windows
	}
}

/* 判断当前系统是否为Linux系统 */
func IsLinux() bool {
	return current_system == Linux_Ubuntu ||
		current_system == Linux_Debian ||
		current_system == Linux_CentOS ||
		current_system == Linux_Unknown
}

/* 判断当前系统是否为Debian发行版本 */
func IsDebian() bool {
	return current_system == Linux_Debian
}

/* 判断当前系统是否为Ubuntu发行版本 */
func IsUbuntu() bool {
	return current_system == Linux_Ubuntu
}

/* 判断当前系统是否为Windows发行版本 */
func IsWindows() bool {
	return current_system == Windows
}

/* 判断当前系统是否为CentOS发行版本 */
func IsCentOS() bool {
	return current_system == Linux_CentOS
}
