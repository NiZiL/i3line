package modules

import (
	"github.com/NiZiL/i3line"
	"os/exec"
)

type TimeModule struct {
}

func (m TimeModule) GetName() string {
	return "time"
}

func (m TimeModule) GenBlock() i3line.Block {
	cmd := exec.Command("date", "+%H:%M:%S")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return i3line.NewErrorBlock(m.GetName(), "local", "error")
	}
	str := string(out[:len(out)-1])
	return i3line.NewDefaultBlock(m.GetName(), "local", str)
}

func (m TimeModule) OnClick(e i3line.Event) {}
