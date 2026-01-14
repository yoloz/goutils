//go:build windows

package cmdutil

import (
	"fmt"
	"os/exec"
	"strings"
	"syscall"
)

// BashExec executes a bash command and returns its output lines.
func BashExec(arg ...string) ([]string, error) {
	return nil, nil
}

// ShExec executes a sh command and returns its output lines.
func ShExec(arg ...string) ([]string, error) {
	return nil, nil
}

// DashExec executes a dash command and returns its output lines.
func DashExec(arg ...string) ([]string, error) {
	return nil, nil
}

// ZshExec executes a zsh command and returns its output lines.
func ZshExec(arg ...string) ([]string, error) {
	return nil, nil
}

// FishExec executes a fish command and returns its output lines.
func FishExec(arg ...string) ([]string, error) {
	return nil, nil
}

// CmdExec executes a command and returns its output lines.
func CmdExec(arg ...string) ([]string, error) {
	return Exec("cmd.exe", "/c", strings.Join(arg, " "))
}

// PowerShellExec executes a PowerShell command and returns its output lines.
func PowerShellExec(arg ...string) ([]string, error) {
	return Exec("powershell.exe", "-Command", strings.Join(arg, " "))
}

// Exec executes a command and returns its output lines.
func Exec(name string, arg ...string) ([]string, error) {
	// Verify executable exists
	if _, err := exec.LookPath(name); err != nil {
		return nil, fmt.Errorf("executable not found: %s", name)
	}
	cmd := exec.Command(name, arg...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		// CmdLine:    fmt.Sprintf(`%s`, strings.Join(arg, " ")),
		HideWindow: true,
	}
	return outputExec(cmd)
}
