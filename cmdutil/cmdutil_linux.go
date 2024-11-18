//go:build linux

package cmdutil

import (
	"os/exec"
)

func GetCmd(cmdExec string) *exec.Cmd {
	return exec.Command("/bin/bash", "-c", cmdExec)
}
