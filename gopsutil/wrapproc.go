package main

//#include <stdlib.h>

import (
	"C"
	"container/list"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/shirou/gopsutil/v3/process"
)

type Process struct {
	Pid        int32    `json:"pid"`
	Name       string   `json:"name"`
	Cmdline    string   `json:"cmdline"`
	CreateTime int64    `json:"createtime"`
	Cwd        string   `json:"cwd"`
	Env        []string `json:"env"`
	Exe        string   `json:"exe"`
	Gids       []int32  `json:"gids"`
	Uids       []int32  `json:"uids"`
	NumFDs     int32    `json:"numfds"`
	NumThreads int32    `json:"numthreads"`
	Ppid       int32    `json:"ppid"`
	Username   string   `json:"username"`
}

func (p Process) String() string {
	s, _ := json.Marshal(p)
	return string(s)
}

func removeSpace(ss []string) []string {
	l := list.New()
	length := 0
	for _, s := range ss {
		if s == "" {
			continue
		}
		l.PushBack(s)
		length++
	}
	sn := make([]string, length)
	i := 0
	for e := l.Front(); e != nil; e = e.Next() {
		sn = append(sn[:i], e.Value.(string))
		i++
	}
	return sn
}

// ProcessList process list
func ProcessList() ([]*Process, error) {
	pss, err := process.Processes()
	if err != nil {
		return nil, err
	}
	l := list.New()
	for i := 0; i < len(pss); i++ {
		p := pss[i]
		// ignore empty cmd
		if li, _ := p.Cmdline(); li == "" {
			continue
		}
		envrion, _ := p.Environ()
		env := removeSpace(envrion)
		name, _ := p.Name()
		cmdline, _ := p.Cmdline()
		createtime, _ := p.CreateTime()
		cwd, _ := p.Cwd()
		exe, _ := p.Exe()
		gids, _ := p.Gids()
		uids, _ := p.Uids()
		numfds, _ := p.NumFDs()
		numthreads, _ := p.NumThreads()
		ppid, _ := p.Ppid()
		username, _ := p.Username()
		l.PushFront(&Process{
			Pid:        p.Pid,
			Name:       name,
			Cmdline:    cmdline,
			CreateTime: createtime,
			Cwd:        cwd,
			Env:        env,
			Exe:        exe,
			Gids:       gids,
			Uids:       uids,
			NumFDs:     numfds,
			NumThreads: numthreads,
			Ppid:       ppid,
			Username:   username,
		})
	}
	var ps []*Process
	for e := l.Front(); e != nil; e = e.Next() {
		ps = append(ps, e.Value.(*Process))
	}
	return ps, nil
}

//export CProcessList
func CProcessList() *C.char {
	ps, err := ProcessList()
	if err != nil {
		fmt.Println("get process list fail..." + err.Error())
		return C.CString("[]")
	}
	bytes, err := json.Marshal(ps)
	if err != nil {
		fmt.Println("marshal process list fail..." + err.Error())
		return C.CString("[]")
	}
	gostr := string(bytes)
	return C.CString(gostr)
}

// ProcessInfo process info
func ProcessByPid(pid int32) (map[string]interface{}, error) {
	proc, err := process.NewProcess(pid)
	if err != nil {
		return nil, err
	}
	createtime, _ := proc.CreateTime()
	cpup, _ := proc.CPUPercent()
	memp, _ := proc.MemoryPercent()
	meminfo, _ := proc.MemoryInfo()

	m := make(map[string]interface{}, 5)
	m["createtime"] = createtime
	m["cpupercent"] = fmt.Sprintf("%.2f", cpup)
	m["mempercent"] = fmt.Sprintf("%.2f", memp)
	m["rss"] = FormatBytes(meminfo.RSS)
	m["vms"] = FormatBytes(meminfo.VMS)
	return m, nil
}

// ProcessInfo process info
func ProcessByCmd(line string) (map[string]interface{}, error) {
	pss, err := process.Processes()
	if err != nil {
		return nil, err
	}
	m := make(map[string]interface{}, 5)
	for _, ps := range pss {
		cmdline, _ := ps.Cmdline()
		if strings.Contains(cmdline, line) {
			createtime, _ := ps.CreateTime()
			cpup, _ := ps.CPUPercent()
			memp, _ := ps.MemoryPercent()
			meminfo, _ := ps.MemoryInfo()
			m["createtime"] = createtime
			m["cpupercent"] = fmt.Sprintf("%.2f", cpup)
			m["mempercent"] = fmt.Sprintf("%.2f", memp)
			m["rss"] = FormatBytes(meminfo.RSS)
			m["vms"] = FormatBytes(meminfo.VMS)
			return m, nil
		}
	}
	return m, nil
}

//export CProcessByCmd
func CProcessByCmd(cmd string) *C.char {
	cards, err := ProcessByCmd(cmd)
	if err != nil {
		fmt.Println("get process info fail..." + err.Error())
		return C.CString("{}")
	}
	bytes, err := json.Marshal(cards)
	if err != nil {
		fmt.Println("marshal process info fail..." + err.Error())
		return C.CString("{}")
	}
	gostr := string(bytes)
	return C.CString(gostr)
}

//export CProcessByPid
func CProcessByPid(pid int32) *C.char {
	cards, err := ProcessByPid(pid)
	if err != nil {
		fmt.Println("get process info fail..." + err.Error())
		return C.CString("{}")
	}
	bytes, err := json.Marshal(cards)
	if err != nil {
		fmt.Println("marshal process info fail..." + err.Error())
		return C.CString("{}")
	}
	gostr := string(bytes)
	return C.CString(gostr)
}
