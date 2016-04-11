package modules

import (
	"github.com/NiZiL/i3line"
	"os/exec"
	"strconv"
	"strings"
)

type SoundModule struct {
}

func (m SoundModule) GetName() string {
	return "sound"
}

func (m SoundModule) GenBlock() i3line.Block {
	cmd := exec.Command("amixer", "sget", "Master")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return i3line.NewErrorBlock(m.GetName(), "local", "error")
	}
	str := string(out[1 : len(out)-2])
	str = str[strings.Index(str, "[")+1:]
	on := strings.Index(str, "on") != -1
	vol := str[:strings.Index(str, "]")-1]
	if on {
		ivol, _ := strconv.Atoi(vol)
		if ivol > 50 {
			str = "🔊 "
		} else if ivol == 0 {
			str = "🔈 "
		} else {
			str = "🔉 "
		}
	} else {
		str = "🔇 "
	}

	return i3line.NewDefaultBlock(m.GetName(), "Master", str+vol+"%")
}

func (m SoundModule) OnClick(e i3line.Event) {
	switch e.Button {
	case 3:
		cmd := exec.Command("amixer", "sset", e.Instance, "toggle")
		cmd.Run()
	case 4:
		cmd := exec.Command("amixer", "sset", e.Instance, "1+")
		cmd.Run()
	case 5:
		cmd := exec.Command("amixer", "sset", e.Instance, "1-")
		cmd.Run()
	}
}
