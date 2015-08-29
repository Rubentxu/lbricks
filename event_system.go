package lbricks
import "github.com/Rubentxu/lbricks/engi"


type EventSystem struct {
	mouseChan    []chan *EventPacket
	keyboardChan []chan *EventPacket
}

func (g *EventSystem) RegisterInputChannel(eventType string, channel chan *EventPacket) {
	switch eventType {
	case "MouseEvent":
		append(g.mouseChan, channel)
	case "KeyboardEvent":
		append(g.keyboardChan, channel)
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
