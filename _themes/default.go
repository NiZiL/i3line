package main

import (
	"github.com/NiZiL/i3line"
	"github.com/NiZiL/i3line/modules"
)

func main() {
	manager := i3line.NewBlockManager()

	manager.Init()
	defer manager.Close()

	manager.AddBlockModule("network", &modules.NetworkModule{})
	manager.AddBlockModule("sound", modules.SoundModule{"Master"})
	manager.AddBlockModule("time", modules.TimeModule{"15:04:05", true})

	manager.Run()
}
