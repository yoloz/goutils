package psutil

//#include <stdlib.h>

import (
	"C"
	"encoding/json"
	"fmt"

	"github.com/shirou/gopsutil/v3/mem"
)

// MemInfo memory info
func MemInfo() (map[string]interface{}, error) {
	mem, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}
	m := make(map[string]interface{}, 4)
	m["total"] = FormatBytes(mem.Total)
	m["available"] = FormatBytes(mem.Available)
	m["used"] = FormatBytes(mem.Used)
	m["usedPercent"] = fmt.Sprintf("%.2f", mem.UsedPercent)
	return m, nil
}

//export CMemInfo
func CMemInfo() *C.char {
	m, err := MemInfo()
	if err != nil {
		fmt.Println("get memory info fail..." + err.Error())
		return C.CString("{}")
	}
	bytes, err := json.Marshal(m)
	if err != nil {
		fmt.Println("marshal memory info fail..." + err.Error())
		return C.CString("{}")
	}
	gostr := string(bytes)
	return C.CString(gostr)
}
