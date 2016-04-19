package i3limod

import (
	"github.com/NiZiL/i3line"
	"net"
)

type NetworkModule struct {
	index int
	len   int
}

func (m *NetworkModule) GenBlock() i3line.Block {
	ifaces, _ := net.Interfaces()
	//remove loopback interface
	for i, iface := range ifaces {
		if iface.Name[0] == 'l' {
			ifaces = append(ifaces[:i], ifaces[i+1:]...)
		}
	}
	m.len = len(ifaces)

	iface := ifaces[m.index]
	addrs, _ := iface.Addrs()
	str := iface.Name
	if len(addrs) > 0 {
		str = str + ": " + addrs[0].String()
		return i3line.NewColorBlock(str, "#00ff00")
	} else {
		str = str + ": down"
		return i3line.NewColorBlock(str, "#ff0000")
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
