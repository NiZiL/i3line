package main

import (
	"github.com/NiZiL/i3line"
	"github.com/NiZiL/i3line/modules"
	"time"
)

func main() {
	manager := i3line.BlockManager{}
	manager.Init()
	defer manager.Close()

	manager.AddModule(&i3limod.NetworkModule{})
	manager.AddModule(i3limod.SoundModule{"Master"})
	manager.AddModule(i3limod.DateModule{"📅 _2/01/2006"})
	manager.AddModule(i3limod.ClockModule{})

	manager.Run(time.Second)
}
