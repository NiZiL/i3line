package main

import (
	"github.com/NiZiL/i3line"
	"github.com/NiZiL/i3line/modules"
)

func main() {
	manager := i3line.NewBlockManager()

	manager.Init()
	defer manager.Close()

	manager.AddBlockModule(&modules.NetworkModule{})
	manager.AddBlockModule(modules.SoundModule{"Master"})
	manager.AddBlockModule(modules.TimeModule{"15:04:05", true})

	manager.Run()
}
