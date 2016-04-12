package main

import (
	"github.com/NiZiL/i3line"
	"github.com/NiZiL/i3line/modules"
)

func main() {
	manager := i3line.NewBlockManager()

	manager.Init()
	defer manager.Close()

	manager.AddBlockModule(modules.SoundModule{"Master"})
	manager.AddBlockModule(modules.TimeModule{"%H:%M:%S", true})

	manager.Run()
}
