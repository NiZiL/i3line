package main

import (
	"github.com/NiZiL/i3line/core"
	"github.com/NiZiL/i3line/modules"
	"time"
)

func main() {
	manager := i3line.BlockManager{}
	manager.Init()
	defer manager.Close()

	manager.AddModule(&modules.NetworkModule{})
	manager.AddModule(modules.SoundModule{"Master"})
	manager.AddModule(modules.DateModule{"📅 _2/01/2006"})
	manager.AddModule(modules.ClockModule{})

	manager.Run(time.Second)
}
