package sysutil

import (
	"testing"
)

func TestGetOSName(t *testing.T) {
	os := GetOSName()
	t.Logf("%v\n", os)
}

func TestGetLinuxRelease(t *testing.T) {
	release := GetLinuxRelease()
	t.Logf("%v\n", release)
}
