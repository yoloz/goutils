package sysuti

import (
	"bufio"
	"os"
	"runtime"
	"strings"
)

// GetOSName get system type
func GetOSName() string {
	return runtime.GOOS
}

// GetLinuxRelease get linux os release
func GetLinuxRelease() (osRelease map[string]string) {
	osRelease = make(map[string]string)
	file, err := os.Open("/etc/os-release")
	if err != nil {
		file, err = os.Open("/etc/*release")
		if err != nil {
			return osRelease
		}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if equalIndex := strings.Index(line, "="); equalIndex >= 0 {
			key, value := line[:equalIndex], line[equalIndex+1:]
			osRelease[key] = strings.Trim(value, `"'`)
		}
	}
	return osRelease
}
