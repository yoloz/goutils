//go:build windows

package cmdutil

import (
	"fmt"
	"os/exec"
	"syscall"
)

func GetCmd(cmdExec string) *exec.Cmd {
	cmd := exec.Command("cmd.exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{CmdLine: fmt.Sprintf(`/c %s`, cmdExec), HideWindow: true}
	return cmd
}
