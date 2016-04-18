package modules

import (
	"github.com/NiZiL/i3line"
	"time"
)

type DateModule struct {
	Format string
	Symbol func() string
}

func (m DateModule) GenBlock() i3line.Block {
	str := m.Symbol() + " " + time.Now().Format(m.Format)
	return i3line.NewDefaultBlock(str)
}

func (m DateModule) OnClick(e i3line.Event) {}

func SyncClockUnicode() string {
	clock0 := [24]string{"🕛", "🕐", "🕑", "🕒", "🕓", "🕔", "🕕", "🕖", "🕗", "🕘", "🕙", "🕚", "🕛", "🕐", "🕑", "🕒", "🕓", "🕔", "🕕", "🕖", "🕗", "🕘", "🕙", "🕚"}
	clock30 := [24]string{"🕧", "🕜", "🕝", "🕞", "🕟", "🕠", "🕡", "🕢", "🕣", "🕤", "🕥", "🕦", "🕧", "🕜", "🕝", "🕞", "🕟", "🕠", "🕡", "🕢", "🕣", "🕤", "🕥", "🕦"}
	t := time.Now()
	if t.Minute() < 30 {
		return clock0[t.Hour()]
	} else {
		return clock30[t.Hour()]
	}
}

func CalendarUnicode() string {
	return "📅"
}
