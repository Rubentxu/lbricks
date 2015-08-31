package lbricks

import (
	"github.com/Rubentxu/lbricks/engi"
	"github.com/Rubentxu/lbricks/goflow"
)


type EventSystem struct {
	stepChan	 []chan *StepEvent
	mouseChan    []chan *MouseEvent
	keyboardChan []chan *KeyboardEvent
}

func (g *EventSystem) RegisterInputChannel(input flow.Port) {
	switch input.Port {
	case "StepEvent" :
		append(g.stepChan, input.Channel)
	case "MouseEvent" :
		append(g.mouseChan, input.Channel)
	case "KeyboardEvent" :
		append(g.keyboardChan, input.Channel)
	}
}

func (g *EventSystem) Preload() {}
func (g *EventSystem) Setup() {}
func (g *EventSystem) Close() {}
func (g *EventSystem) Update(dt float32) {}
func (g *EventSystem) Step(step float64, numStep uint32) {}
func (g *EventSystem) Render() {}
func (g *EventSystem) Resize(w, h int) {}

func (g *EventSystem) Mouse(x, y float32, action engi.MouseAction) {
	EventPacket := &EventPacket{1, &MouseEvent{x, y, 0.0, action}}
	for _, e := range g.mouseChan {
		e <- EventPacket
	}
}

func (g *EventSystem) Scroll(amount float32) {}

func (g *EventSystem) Key(key engi.Key, modifier engi.Modifier, action engi.KeyAction) {
	if key == engi.Escape {
		engi.Exit()
	}
	EventPacket := &EventPacket{1, &KeyboardEvent{key, modifier, action}}
	for _, e := range g.keyboardChan {
		e <- EventPacket
	}
}

func (g *EventSystem) Type(char rune) {}
