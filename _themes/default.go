package main

import (
	"github.com/NiZiL/i3line"
	"github.com/NiZiL/i3line/modules"
	"time"
)

func main() {
	manager := i3line.NewBlockManager()

	manager.Init()
	defer manager.Close()

	manager.AddBlockModule(&i3limod.NetworkModule{})
	manager.AddBlockModule(i3limod.SoundModule{"Master"})
	manager.AddBlockModule(i3limod.DateModule{"📅 _2/01/2006"})
	manager.AddBlockModule(i3limod.ClockModule{})

	manager.Run(time.Second)
}
