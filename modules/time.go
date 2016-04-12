package modules

import (
	"github.com/NiZiL/i3line"
	"time"
)

type TimeModule struct {
	Format string
	Clock  bool
}

func (m TimeModule) GetName() string {
	return "time"
}

func (m TimeModule) GenBlock() i3line.Block {
	t := time.Now()
	str := t.Format(m.Format)
	if m.Clock {
		str = clockUnicode(t) + " " + str
	}
	return i3line.NewDefaultBlock(m.GetName(), "local", str)
}

func (m TimeModule) OnClick(e i3line.Event) {}

func clockUnicode(t time.Time) string {
	clock0 := [24]string{"🕛", "🕐", "🕑", "🕒", "🕓", "🕔", "🕕", "🕖", "🕗", "🕘", "🕙", "🕚", "🕛", "🕐", "🕑", "🕒", "🕓", "🕔", "🕕", "🕖", "🕗", "🕘", "🕙", "🕚"}
	clock30 := [24]string{"🕧", "🕜", "🕝", "🕞", "🕟", "🕠", "🕡", "🕢", "🕣", "🕤", "🕥", "🕦", "🕧", "🕜", "🕝", "🕞", "🕟", "🕠", "🕡", "🕢", "🕣", "🕤", "🕥", "🕦"}
	if t.Minute() < 30 {
		return clock0[t.Hour()]
	} else {
		return clock30[t.Hour()]
	}
}
