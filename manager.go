package i3line

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

type BlockModule interface {
	GenBlock() Block
	OnClick(Event) bool
}

type BlockManager struct {
	modules []BlockModule
	blocks  []Block
}

func NewBlockManager() *BlockManager {
	manager := new(BlockManager)
	manager.modules = make([]BlockModule, 0)
	return manager
}

func (m *BlockManager) Init() {
	fmt.Println(`{ "version": 1, "click_events": true }`)
	fmt.Println(`[`)
}

func (m *BlockManager) Close() {
	fmt.Println(`]`)
}

func (m *BlockManager) AddBlockModule(module BlockModule) {
	m.modules = append(m.modules, module)
}

func (m *BlockManager) Run(refreshRate time.Duration) {
	// init internal blocks buffer
	m.blocks = make([]Block, len(m.modules))

	m.updateAllBlocks()
	m.sync()
	// a simple ticker, will be removed soon
	// we need something to provide a different refresh rate by modules
	c := time.Tick(refreshRate)
	go func() {
		for {
			select {
			case <-c:
				m.updateAllBlocks()
				m.sync()
			}
		}
	}()
	m.listenEvent()
}

func (m *BlockManager) updateAllBlocks() {
	for id, module := range m.modules {
		b := module.GenBlock()
		b.Name = strconv.Itoa(id)
		b.Instance = strconv.Itoa(id)
		m.blocks[id] = b
	}
}

func (m *BlockManager) updateBlock(id int) {
	b := m.modules[id].GenBlock()
	b.Name = strconv.Itoa(id)
	b.Instance = strconv.Itoa(id)
	m.blocks[id] = b
}

func (m *BlockManager) sync() {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(m.blocks)
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
		go m.handleEvent(event)
	}

	// read closing bracket
	if _, err := decoder.Token(); err != nil {
		panic(err)
	}
}

func (m *BlockManager) handleEvent(e Event) {
	id, _ := strconv.Atoi(e.Name)
	if m.modules[id].OnClick(e) {
		m.updateBlock(id)
		m.sync()
	}
}
