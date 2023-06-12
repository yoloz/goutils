package psutil

//#include <stdlib.h>

import (
	"C"
	"encoding/json"
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/load"
)

// CpuInfo cpu info
func CpuInfo() (map[string]interface{}, error) {
	physical, err := cpu.Counts(false)
	if err != nil {
		return nil, err
	}
	logical, err := cpu.Counts(true)
	if err != nil {
		return nil, err
	}
	infoStats, err := cpu.Info()
	if err != nil {
		return nil, err
	}
	m := make(map[string]interface{}, 7)
	m["physicalCount"] = physical
	m["logicalCount"] = logical
	m["vendorId"] = infoStats[0].VendorID
	m["physicalId"] = infoStats[0].PhysicalID
	m["modelName"] = infoStats[0].ModelName
	m["family"] = infoStats[0].Family
	m["mhz"] = infoStats[0].Mhz
	return m, nil
}

// CpuStat cpu status
func CpuStat() (map[string]interface{}, error) {
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return nil, err
	}
	stats, err := cpu.Times(false)
	if err != nil {
		return nil, err
	}
	avg, err := load.Avg()
	if err != nil {
		return nil, err
	}
	load1 := fmt.Sprintf("%.1f", avg.Load1)
	load5 := fmt.Sprintf("%.1f", avg.Load5)
	load15 := fmt.Sprintf("%.1f", avg.Load15)
	m := make(map[string]interface{}, 10)
	m["percent"] = fmt.Sprintf("%.1f", percent[0])
	m["user"] = fmt.Sprintf("%.1f", stats[0].User)
	m["nice"] = fmt.Sprintf("%.1f", stats[0].Nice)
	m["system"] = fmt.Sprintf("%.1f", stats[0].System)
	m["idle"] = fmt.Sprintf("%.1f", stats[0].Idle)
	m["iowait"] = fmt.Sprintf("%.1f", stats[0].Iowait)
	m["irq"] = fmt.Sprintf("%.1f", stats[0].Irq)
	m["softirq"] = fmt.Sprintf("%.1f", stats[0].Softirq)
	m["steal"] = fmt.Sprintf("%.1f", stats[0].Steal)
	m["load"] = fmt.Sprintf("%s %s %s", load1, load5, load15)
	return m, nil
}

//export CCpuInfo
func CCpuInfo() *C.char {
	m, err := CpuInfo()
	if err != nil {
		fmt.Println("get cpu info fail..." + err.Error())
		return C.CString("{}")
	}
	bytes, err := json.Marshal(m)
	if err != nil {
		fmt.Println("marshal cpu info fail..." + err.Error())
		return C.CString("{}")
	}
	gostr := string(bytes)
	return C.CString(gostr)
}

//export CCpuStat
func CCpuStat() *C.char {
	m, err := CpuStat()
	if err != nil {
		fmt.Println("get cpu stat fail..." + err.Error())
		return C.CString("{}")
	}
	bytes, err := json.Marshal(m)
	if err != nil {
		fmt.Println("marshal cpu stat fail..." + err.Error())
		return C.CString("{}")
	}
	gostr := string(bytes)
	return C.CString(gostr)
}
