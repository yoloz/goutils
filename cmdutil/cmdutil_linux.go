//go:build linux

package cmdutil

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// CmdExec executes a command and returns its output lines.
func CmdExec(arg ...string) ([]string, error) {
	return nil, nil
}

// PowerShellExec executes a PowerShell command and returns its output lines.
func PowerShellExec(arg ...string) ([]string, error) {
	return nil, nil
}

// BashExec executes a bash command and returns its output lines.
func BashExec(arg ...string) ([]string, error) {
	return Exec("/bin/bash", arg...)
}

// ShExec executes a sh command and returns its output lines.
func ShExec(arg ...string) ([]string, error) {
	return Exec("/bin/sh", arg...)
}

// DashExec executes a dash command and returns its output lines.
func DashExec(arg ...string) ([]string, error) {
	return Exec("/bin/dash", arg...)
}

// ZshExec executes a zsh command and returns its output lines.
func ZshExec(arg ...string) ([]string, error) {
	return Exec("/bin/zsh", arg...)
}

// FishExec executes a fish command and returns its output lines.
func FishExec(arg ...string) ([]string, error) {
	return Exec("/bin/fish", arg...)
}

// Exec executes a command and returns its output lines.
func Exec(name string, arg ...string) ([]string, error) {
	if err := validateShell(name); err != nil {
		return nil, err
	}

	cmd := exec.Command(name, "-c", strings.Join(arg, " "))
	return outputExec(cmd)
}

// Validate shell existence before execution
func validateShell(shellPath string) error {
	if _, err := os.Stat(shellPath); os.IsNotExist(err) {
		return fmt.Errorf("shell does not exist: %s", shellPath)
	}
	return nil
}
