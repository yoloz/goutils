package cmdutil

import (
	"errors"
	"fmt"
	"os/exec"
	"syscall"

	"github.com/yoloz/goutils/sysutil"
)

func GetCmd(cmdExec string) (*exec.Cmd, error) {
	switch os := sysutil.GetOSName(); os {
	case "windows":
		cmd := exec.Command("cmd.exe")
		cmd.SysProcAttr = &syscall.SysProcAttr{CmdLine: fmt.Sprintf(`/c %s`, cmdExec), HideWindow: true}
		return cmd, nil
	case "linux":
		return exec.Command("/bin/bash", "-c", cmdExec), nil
	default:
		return nil, errors.New(os + " cmd not support.")
	}
}
