package psutil

//#include <stdlib.h>

import (
	"C"

	"container/list"
	"encoding/json"
	"fmt"

	"github.com/shirou/gopsutil/v3/net"
)

type NetCard struct {
	MTU          int                   `json:"mtu"`
	Name         string                `json:"name"`
	HardwareAddr string                `json:"hardwareAddr"`
	Flags        []string              `json:"flags"`
	Addrs        net.InterfaceAddrList `json:"addrs"`
}

func (c NetCard) String() string {
	s, _ := json.Marshal(c)
	return string(s)
}

// isUp check flag is not null and is up
func isUp(flags []string) bool {
	if len(flags) == 0 {
		return false
	}
	for _, f := range flags {
		if f == "up" {
			return true
		}
	}
	return false
}

// NCList net card list
func NCList() ([]*NetCard, error) {
	faces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	var l list.List
	length := 0
	for i := 0; i < len(faces); i++ {
		card := faces[i]
		// ignore 127.0.0.1
		if card.HardwareAddr == "" {
			continue
		}
		// ignore down card
		if !isUp(card.Flags) {
			continue
		}
		l.PushFront(&NetCard{
			Name:         card.Name,
			HardwareAddr: card.HardwareAddr,
			Flags:        card.Flags,
			MTU:          card.MTU,
			Addrs:        card.Addrs,
		})
		length++
	}
	var cards []*NetCard = make([]*NetCard, length)
	i := 0
	for e := l.Front(); e != nil; e = e.Next() {
		cards = append(cards[:i], e.Value.(*NetCard))
		i++
	}
	return cards, nil
}

//export CNCList
func CNCList() *C.char {
	cards, err := NCList()
	if err != nil {
		fmt.Println("get netCard info fail..." + err.Error())
		return C.CString("[]")
	}
	bytes, err := json.Marshal(cards)
	if err != nil {
		fmt.Println("marshal netCard info fail..." + err.Error())
		return C.CString("[]")
	}
	gostr := string(bytes)
	return C.CString(gostr)
}
