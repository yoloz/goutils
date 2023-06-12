package psutil

//#include <stdlib.h>

import (
	"C"
	"encoding/json"
	"fmt"

	"github.com/shirou/gopsutil/v3/host"
)

// HostInfo host info
func HostInfo() (*map[string]interface{}, error) {
	host, err := host.Info()
	if err != nil {
		return nil, err
	}
	m := make(map[string]interface{}, 13)
	m["hostname"] = host.Hostname
	m["hostid"] = host.HostID
	m["uptime"] = host.Uptime
	m["boottime"] = host.BootTime
	m["procs"] = host.Procs
	m["os"] = host.OS
	m["platform"] = host.Platform
	m["platformFamily"] = host.PlatformFamily
	m["platformVersion"] = host.PlatformVersion
	m["kernelArch"] = host.KernelArch
	m["kernelVersion"] = host.KernelVersion
	m["virtualizationSystem"] = host.VirtualizationSystem
	m["virtualizationRole"] = host.VirtualizationRole
	return &m, nil
}

//export CHostInfo
func CHostInfo() *C.char {
	m, err := HostInfo()
	if err != nil {
		fmt.Println("get host info fail..." + err.Error())
		return C.CString("{}")
	}
	bytes, err := json.Marshal(m)
	if err != nil {
		fmt.Println("marshal host info fail..." + err.Error())
		return C.CString("{}")
	}
	gostr := string(bytes)
	return C.CString(gostr)
}
