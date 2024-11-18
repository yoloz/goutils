package cmdutil

import (
	"testing"
)

func TestLiuxCMD(t *testing.T) {
	cmd := GetCmd("echo ${PATH}")
	bytes, err := cmd.Output()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s\n", string(bytes))
}

func TestWindowCMD(t *testing.T) {
	cmd := GetCmd("set path")
	bytes, err := cmd.Output()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s\n", string(bytes))
}
