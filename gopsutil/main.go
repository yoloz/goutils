package main

import (
	"container/list"
	"encoding/json"
	"fmt"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/shirou/gopsutil/v3/process"
)

func main() {
	// hostInfo()
	// cpuInfo()
	// memInfo()
	// diskInfo()
	// netInfo()
	procInfo()

}

func hostInfo() {
	fmt.Println("主机信息：")
	host, _ := host.Info()
	fmt.Printf(" %v\n", host)
}

func cpuInfo() {
	fmt.Println("CPU信息：")
	cpu, _ := cpu.Info()
	fmt.Printf(" %v\n", cpu)
}

func memInfo() {
	fmt.Println("内存信息：")
	mem, _ := mem.VirtualMemory()
	fmt.Printf(" %v\n", mem)
}

func diskInfo() {
	fmt.Println("磁盘信息：")
	//io counters 暂时不需要
	// dev, _ := disk.IOCounters()
	// fmt.Printf(" %v\n", dev)
	fmt.Println("分区信息：")
	part, _ := disk.Partitions(false)
	fmt.Printf(" %v\n", part)
}

func netInfo() {
	fmt.Println("网络信息：")
	network, _ := net.Interfaces()
	fmt.Printf(" %v\n", network)
}

func procInfo() {
	fmt.Println("进程信息：")
	var l list.List
	ps, _ := process.Processes()
	for _, p := range ps {
		m := make(map[string]interface{}, 13)
		m["pid"] = p.Pid
		m["name"], _ = p.Name()
		m["cmdline"], _ = p.Cmdline()
		m["createtime"], _ = p.CreateTime()
		m["cwd"], _ = p.Cwd()
		m["env"], _ = p.Environ()
		m["exe"], _ = p.Exe()
		m["gids"], _ = p.Gids()
		m["uids"], _ = p.Uids()
		m["numfds"], _ = p.NumFDs()
		m["numthreads"], _ = p.NumThreads()
		m["ppid"], _ = p.Ppid()
		m["username"], _ = p.Username()
		l.PushBack(m)
	}
	fmt.Printf("[")
	for i := l.Front(); i != nil; i = i.Next() {
		pss, _ := json.Marshal(i.Value)
		fmt.Printf("%s,", string(pss))
	}
	fmt.Printf("]")
}
