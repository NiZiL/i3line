package i3line

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type BlockModule interface {
	GenBlock() Block
	OnClick(Event) bool
}

type BlockManager struct {
	modules map[string]BlockModule
	order   []string
}

func NewBlockManager() *BlockManager {
	manager := new(BlockManager)
	manager.modules = make(map[string]BlockModule)
	manager.order = make([]string, 0)
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

func (m *BlockManager) Run(refreshRate time.Duration) {
	m.updateBlocks()
	c := time.Tick(refreshRate)
	go func() {
		for {
			select {
			case <-c:
				m.updateBlocks()
			}
		}
	}()
	m.listenEvent()
}

func (m *BlockManager) updateBlocks() {
	blocks := make([]Block, len(m.order))
	for i, modName := range m.order {
		b := m.modules[modName].GenBlock()
		b.Name = modName
		b.Instance = modName
		blocks[i] = b
	}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(blocks)
	fmt.Println(buf.String() + ",")
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

func (m *BlockManager) handleEvent(e Event) {
	if m.modules[e.Name].OnClick(e) {
		m.updateBlocks()
	}
}
