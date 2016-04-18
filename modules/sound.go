package modules

import (
	"github.com/NiZiL/i3line"
	"os/exec"
	"strconv"
	"strings"
)

type SoundModule struct {
	Channel string
}

func (m SoundModule) GenBlock() i3line.Block {
	cmd := exec.Command("amixer", "sget", m.Channel)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return i3line.NewErrorBlock()
	}
	str := string(out)
	on := strings.Index(str[len(str)-5:], "on") != -1
	vol := str[strings.Index(str, "[")+1 : strings.Index(str, "]")-1]
	if on {
		ivol, _ := strconv.Atoi(vol)
		if ivol > 80 {
			str = "🔊 "
		} else if ivol == 0 {
			str = "🔈 "
		} else {
			str = "🔉 "
		}
	} else {
		str = "🔇 "
	}

	return i3line.NewDefaultBlock(str + vol + "%")
}

func (m SoundModule) OnClick(e i3line.Event) {
	switch e.Button {
	case 3:
		cmd := exec.Command("amixer", "sset", m.Channel, "toggle")
		cmd.Run()
	case 4:
		cmd := exec.Command("amixer", "sset", m.Channel, "1+")
		cmd.Run()
	case 5:
		cmd := exec.Command("amixer", "sset", m.Channel, "1-")
		cmd.Run()
	}
}
