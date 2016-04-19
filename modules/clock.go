package modules

import (
	"github.com/NiZiL/i3line"
	"time"
)

type ClockModule struct {
}

func (m ClockModule) GenBlock() i3line.Block {
	t := time.Now()
	str := clockUnicode(t) + " " + t.Format("15:04:05")
	return i3line.NewDefaultBlock(str)
}

func (m ClockModule) OnClick(e i3line.Event) bool {
	return false
}

func clockUnicode(t time.Time) string {
	clock0 := [24]string{"🕛", "🕐", "🕑", "🕒", "🕓", "🕔", "🕕", "🕖", "🕗", "🕘", "🕙", "🕚", "🕛", "🕐", "🕑", "🕒", "🕓", "🕔", "🕕", "🕖", "🕗", "🕘", "🕙", "🕚"}
	clock30 := [24]string{"🕧", "🕜", "🕝", "🕞", "🕟", "🕠", "🕡", "🕢", "🕣", "🕤", "🕥", "🕦", "🕧", "🕜", "🕝", "🕞", "🕟", "🕠", "🕡", "🕢", "🕣", "🕤", "🕥", "🕦"}
	if t.Minute() < 30 {
		return clock0[t.Hour()]
	} else {
		return clock30[t.Hour()]
	}
}
