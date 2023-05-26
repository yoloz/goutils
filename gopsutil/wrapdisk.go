package main

//#include <stdlib.h>

import (
	"C"
	"fmt"

	"github.com/shirou/gopsutil/v3/disk"
)
import "encoding/json"

type Partition struct {
	Device      string `json:"device"`
	Mountpoint  string `json:"mountpoint"`
	Fstype      string `json:"fstype"`
	Total       string `json:"total"`
	Free        string `json:"free"`
	Used        string `json:"used"`
	UsedPercent string `json:"usedPercent"`
}

// PartitionInfo partition info
func PartitionInfo() (map[string]interface{}, error) {
	partions, err := disk.Partitions(false)
	if err != nil {
		return nil, err
	}
	parts := make([]*Partition, len(partions))
	var total, free, used uint64
	for i := 0; i < len(partions); i++ {
		stat := partions[i]
		usage, err := disk.Usage(stat.Mountpoint)
		if err != nil {
			return nil, err
		}
		total += usage.Total
		free += usage.Free
		used += usage.Used
		parts[i] = &Partition{
			Device:      stat.Device,
			Mountpoint:  stat.Mountpoint,
			Fstype:      stat.Fstype,
			Total:       FormatBytes(usage.Total),
			Free:        FormatBytes(usage.Free),
			Used:        FormatBytes(usage.Used),
			UsedPercent: fmt.Sprintf("%.2f", usage.UsedPercent),
		}
	}
	usedPercent := (float64(used) / float64(used+free)) * 100.0
	m := make(map[string]interface{}, 5)
	m["total"] = FormatBytes(total)
	m["used"] = FormatBytes(used)
	m["free"] = FormatBytes(free)
	m["usedPercent"] = fmt.Sprintf("%.2f", usedPercent)
	m["list"] = parts
	return m, nil
}

func (p Partition) String() string {
	s, _ := json.Marshal(p)
	return string(s)
}

//export CPartitionInfo
func CPartitionInfo() *C.char {
	m, err := PartitionInfo()
	if err != nil {
		fmt.Println("get disk partion fail..." + err.Error())
		return C.CString("{}")
	}
	bytes, err := json.Marshal(m)
	if err != nil {
		fmt.Println("marshal disk partion fail..." + err.Error())
		return C.CString("{}")
	}
	gostr := string(bytes)
	return C.CString(gostr)
}
