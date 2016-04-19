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

	manager.AddBlockModule("network", &modules.NetworkModule{})
	manager.AddBlockModule("sound", modules.SoundModule{"Master"})
	manager.AddBlockModule("date", modules.DateModule{"📅 _2/01/2006"})
	manager.AddBlockModule("clock", modules.ClockModule{})

	manager.Run(time.Second)
}
