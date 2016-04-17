package i3line

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

type BlockModule interface {
	GenBlock() Block
	OnClick(Event)
}

type BlockManager struct {
	modules  map[string]BlockModule
	order    []string
	lastSend string
}

func NewBlockManager() *BlockManager {
	manager := new(BlockManager)
	manager.modules = make(map[string]BlockModule)
	manager.order = make([]string, 0)
	manager.lastSend = ""
	return manager
}

func (m *BlockManager) Init() {
	fmt.Println(`{ "version": 1, "click_events": true }`)
	fmt.Println(`[`)
}

func (m *BlockManager) Close() {
	fmt.Println(`]`)
}

func (m *BlockManager) AddBlockModule(name string, module BlockModule) {
	m.order = append(m.order, name)
	m.modules[name] = module
}

func (m *BlockManager) Run() {
	go func() {
		//TODO scheduler instead of for loop
		for {
			var blocks []Block
			var block Block
			for _, name := range m.order {
				block = m.modules[name].GenBlock()
				block.Name = name
				block.Instance = name
				blocks = append(blocks, block)
			}
			m.refreshBlocks(blocks)
		}
	}()
	m.listenEvent()
}

func (m *BlockManager) refreshBlocks(blocks []Block) {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(blocks)
	if buf.String() != m.lastSend {
		fmt.Println(buf.String() + ",")
		m.lastSend = buf.String()
	}
}

type Event struct {
	Name     string `json: "name"`
	Instance string `json: "instance"`
	Button   int    `json: "button"`
	X        int    `json: "x"`
	Y        int    `json: "y"`
}

func (m *BlockManager) listenEvent() {
	// new json decoder from stdin
	decoder := json.NewDecoder(bufio.NewReader(os.Stdin))

	// read opening bracket
	if _, err := decoder.Token(); err != nil {
		panic(err)
	}

	// read while i3bar sends clickEvent, should be infinite
	var event Event
	for decoder.More() {
		err := decoder.Decode(&event)
		if err != nil {
			panic(err)
		}
		go m.modules[event.Name].OnClick(event)
	}

	// read closing bracket
	if _, err := decoder.Token(); err != nil {
		panic(err)
	}
}
