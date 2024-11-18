package pathutil

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// GetCurrentPath get current directoy
func GetAppPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return "", errors.New(`error: Absolute Path format Fail`)
	}
	return string(path[0 : i+1]), nil
}
