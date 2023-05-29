package main

import (
	"fmt"
	"testing"
)

func TestHostInfo(t *testing.T) {
	m, _ := HostInfo()
	fmt.Printf("%v\n", m)
}

func TestCpuInfo(t *testing.T) {
	m, _ := CpuInfo()
	fmt.Printf("%v\n", m)

}

func TestCpuStat(t *testing.T) {
	m, _ := CpuStat()
	fmt.Printf("%v\n", m)
}

func TestMemInfo(t *testing.T) {
	m, _ := MemInfo()
	fmt.Printf("%v\n", m)
}

func TestDiskInfo(t *testing.T) {
	m, _ := PartitionInfo()
	fmt.Printf("%v\n", m)
}

func TestNCList(t *testing.T) {
	m, _ := NCList()
	fmt.Printf("%v\n", m)
}

func TestProcInfo(t *testing.T) {
	m, _ := ProcessList()
	fmt.Printf("%v\n", m)
}

func TestProcessByPid(t *testing.T) {
	m, _ := ProcessByPid(1247)
	fmt.Printf("%v\n", m)
}

func TestProcessByCmd(t *testing.T) {
	m, _ := ProcessByCmd("code")
	fmt.Printf("%v\n", m)
}
