package modules

import (
	"github.com/NiZiL/i3line"
	"os/exec"
	"strconv"
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
	return i3line.NewDefaultBlock(m.GetName(), "local", clockUnicode(str)+" "+str)
}

func (m TimeModule) OnClick(e i3line.Event) {}

func clockUnicode(str string) string {
	h, _ := strconv.Atoi(str[0:2])
	m, _ := strconv.Atoi(str[3:5])
	switch h {
	case 0:
		fallthrough
	case 12:
		if m < 30 {
			return "🕛"
		} else {
			return "🕧"
		}

	case 1:
		fallthrough
	case 13:
		if m < 30 {
			return "🕐"
		} else {
			return "🕜"
		}

	case 2:
		fallthrough
	case 14:
		if m < 30 {
			return "🕑"
		} else {
			return "🕝"
		}

	case 3:
		fallthrough
	case 15:
		if m < 30 {
			return "🕒"
		} else {
			return "🕞"
		}

	case 4:
		fallthrough
	case 16:
		if m < 30 {
			return "🕓"
		} else {
			return "🕟"
		}

	case 5:
		fallthrough
	case 17:
		if m < 30 {
			return "🕔"
		} else {
			return "🕠"
		}

	case 6:
		fallthrough
	case 18:
		if m < 30 {
			return "🕕"
		} else {
			return "🕡"
		}

	case 7:
		fallthrough
	case 19:
		if m < 30 {
			return "🕖"
		} else {
			return "🕢"
		}

	case 8:
		fallthrough
	case 20:
		if m < 30 {
			return "🕗"
		} else {
			return "🕣"
		}

	case 9:
		fallthrough
	case 21:
		if m < 30 {
			return "🕘"
		} else {
			return "🕤"
		}

	case 10:
		fallthrough
	case 22:
		if m < 30 {
			return "🕙"
		} else {
			return "🕥"
		}

	case 11:
		fallthrough
	case 23:
		if m < 30 {
			return "🕚"
		} else {
			return "🕦"
		}

	default:
		return "⌚"
	}
}
