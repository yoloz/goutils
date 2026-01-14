package cmdutil

import (
	"testing"
)

func TestBashExec(t *testing.T) {
	lines, err := BashExec("ip addr show", `| awk '/^[0-9]+:/ {gsub(/:/, "", $2);interface = $2;next;}/link\/ether/ {mac = $2;}/inet / {ip = $2;split(ip, parts, "/");printf "%-s %-s %-s\n", interface, parts[1], mac;}'`)
	if err != nil {
		t.Fatal(err)
	}
	for _, line := range lines {
		t.Logf("%s\n", line)
	}
}

func TestShExec(t *testing.T) {
	lines, err := ShExec("ip addr show", `| awk '/^[0-9]+:/ {gsub(/:/, "", $2);interface = $2;next;}/link\/ether/ {mac = $2;}/inet / {ip = $2;split(ip, parts, "/");printf "%-s %-s %-s\n", interface, parts[1], mac;}'`)
	if err != nil {
		t.Fatal(err)
	}
	for _, line := range lines {
		t.Logf("%s\n", line)
	}
}
func TestCmdExec(t *testing.T) {
	lines, err := CmdExec("netsh", "interface", "show", "interface")
	if err != nil {
		t.Fatal(err)
	}
	for _, line := range lines {
		t.Logf("%s\n", line)
	}
}

func TestPowerShellExec(t *testing.T) {
	lines, err := PowerShellExec("Get-NetAdapter | Where-Object {$_.Status -eq 'Up'} | Select-Object Name, InterfaceDescription, MacAddress | ConvertTo-Csv")
	if err != nil {
		t.Fatal(err)
	}
	for _, line := range lines {
		t.Logf("%s\n", line)
	}
}
