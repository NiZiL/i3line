package modules

import (
	"github.com/NiZiL/i3line"
	"time"
)

type DateModule struct {
	Format string
}

func (m DateModule) GenBlock() i3line.Block {
	return i3line.NewDefaultBlock(time.Now().Format(m.Format))
}

func (m DateModule) OnClick(e i3line.Event) bool {
	return false
}
