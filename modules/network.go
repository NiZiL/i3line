package i3limod

import (
	"github.com/NiZiL/i3line"
	"net"
	"strings"
)

type NetworkModule struct {
	index int
	len   int
}

func (m *NetworkModule) GenBlock() i3line.Block {
	ifaces, _ := net.Interfaces()
	//remove loopback interface
	for i, iface := range ifaces {
		if strings.Contains(iface.Flags.String(), net.FlagLoopback.String()) {
			ifaces = append(ifaces[:i], ifaces[i+1:]...)
		}
	}
	m.len = len(ifaces)

	iface := ifaces[m.index]
	addrs, _ := iface.Addrs()
	str := "<span font_size='x-small'>" + iface.Name + ": </span>"
	if len(addrs) > 0 {
		ip := addrs[0].String()
		str = str + "<span foreground='green'>" + ip[:len(ip)-3] + "</span><span foreground='grey'>" + ip[len(ip)-3:] + "</span>"
		return i3line.NewPangoBlock(str)
	} else {
		str = str + "<span foreground='red'>down</span>"
		return i3line.NewPangoBlock(str)
	}
}

func (m *NetworkModule) OnClick(e i3line.Event) bool {
	res := false
	switch e.Button {
	case 1:
		fallthrough
	case 4:
		m.index = m.index + 1
		res = true
	case 3:
		fallthrough
	case 5:
		m.index = m.index - 1
		res = true
	}
	if m.index < 0 {
		m.index = m.len - 1
	} else if m.index >= m.len {
		m.index = 0
	}
	return res
}
