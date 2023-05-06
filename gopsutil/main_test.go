package main

import "testing"

func TestHostInfo(t *testing.T) {
	hostInfo()
}

func TestCpuInfo(t *testing.T) {
	cpuInfo()
}

func TestMemInfo(t *testing.T) {
	memInfo()
}

func TestDiskInfo(t *testing.T) {
	diskInfo()
}

func TestNetInfo(t *testing.T) {
	netInfo()
}

func TestProcInfo(t *testing.T) {
	procInfo()
}
