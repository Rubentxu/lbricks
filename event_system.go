package lbricks

import (
	"github.com/Rubentxu/lbricks/engi"
	"github.com/Rubentxu/lbricks/goflow"
)


type EventSystem struct {
	preloadEventChannnels 	[]chan *PreloadEvent
	setupEventChannels	 	[]chan *SetupEvent
	closeEventChannels	 	[]chan *CloseEvent
	updateEventChannels	 	[]chan *UpdateEvent
	renderEventChannels	 	[]chan *RenderEvent
	resizeEventChannels	 	[]chan *ResizeEvent
	stepEventChannels 		[]chan *StepEvent
	mouseEventChannels 		[]chan *MouseEvent
	scrollEventChannels	 	[]chan *ScrollEvent
	keyEventChannels 		[]chan *KeyEvent
	typeKeyEventChannels 	[]chan *TypeKeyEvent
}

func CreateEventSystem(capacity int) *EventSystem  {
	eventSystem := new(EventSystem)
	eventSystem.preloadEventChannnels = make([]chan *PreloadEvent, 50, capacity)
	eventSystem.setupEventChannels =	make([]chan *SetupEvent, 50, capacity)
	eventSystem.closeEventChannels = 	make([]chan *CloseEvent, 50, capacity)
	eventSystem.updateEventChannels = 	make([]chan *UpdateEvent, 50, capacity)
	eventSystem.renderEventChannels = 	make([]chan *RenderEvent, 50, capacity)
	eventSystem.resizeEventChannels = 	make([]chan *ResizeEvent, 50, capacity)
	eventSystem.stepEventChannels = 	make([]chan *StepEvent, 50, capacity)
	eventSystem.mouseEventChannels = 	make([]chan *MouseEvent, 50, capacity)
	eventSystem.scrollEventChannels = 	make([]chan *ScrollEvent, 50, capacity)
	eventSystem.keyEventChannels = 		make([]chan *KeyEvent, 50, capacity)
	eventSystem.typeKeyEventChannels = 	make([]chan *TypeKeyEvent, 50, capacity)
	return eventSystem
}

func (g *EventSystem) RegisterInputChannel(input flow.Port) {
	switch input.Port {
	case "PreloadEvent" :
		append(g.preloadEventChannnels, input.Channel.(PreloadEvent))
	case "SetupEvent" :
		append(g.setupEventChannels, input.Channel.(SetupEvent))
	case "CloseEvent" :
		append(g.closeEventChannels, input.Channel.(CloseEvent))
	case "UpdateEvent" :
		append(g.updateEventChannels, input.Channel.(UpdateEvent))
	case "RenderEvent" :
		append(g.renderEventChannels, input.Channel.(RenderEvent))
	case "ResizeEvent" :
		append(g.resizeEventChannels, input.Channel.(ResizeEvent))
	case "StepEvent" :
		append(g.stepEventChannels, input.Channel.(StepEvent))
	case "MouseEvent" :
		append(g.mouseEventChannels, input.Channel.(MouseEvent))
	case "ScrollEvent" :
		append(g.scrollEventChannels, input.Channel.(ScrollEvent))
	case "KeyEvent" :
		append(g.keyEventChannels, input.Channel.(KeyEvent))
	case "TypeKeyEvent" :
		append(g.typeKeyEventChannels, input.Channel.(TypeKeyEvent))
	}
}

func (g *EventSystem) Preload() {
	event := &PreloadEvent{}
	for _, e := range g.preloadEventChannnels {
		e <- event
	}
}

func (g *EventSystem) Setup() {
	event := &SetupEvent{}
	for _, e := range g.setupEventChannels {
		e <- event
	}
}

func (g *EventSystem) Close() {
	event := &CloseEvent{}
	for _, e := range g.closeEventChannels {
		e <- event
	}
}

func (g *EventSystem) Update(dt float32) {
	event := &UpdateEvent{dt}
	for _, e := range g.updateEventChannels {
		e <- event
	}
}

func (g *EventSystem) Step(step float64, numStep uint32) {
	event := &StepEvent{step, numStep}
	for _, e := range g.stepEventChannels {
		e <- event
	}
}

func (g *EventSystem) Render() {
	event := &RenderEvent{}
	for _, e := range g.renderEventChannels {
		e <- event
	}
}

func (g *EventSystem) Resize(w, h int) {
	event := &ResizeEvent{w, h}
	for _, e := range g.resizeEventChannels {
		e <- event
	}
}

func (g *EventSystem) Mouse(x, y float32, action engi.MouseAction) {
	event := &MouseEvent{x, y, 0.0, action}
	for _, e := range g.mouseEventChannels {
		e <- event
	}
}

func (g *EventSystem) Scroll(amount float32) {
	var action engi.MouseAction = engi.WHEEL_UP
	if amount < 0 {
		action = engi.WHEEL_DOWN
	}
	event := &MouseEvent{0, 0, amount, action}
	for _, e := range g.scrollEventChannels {
		e <- event
	}
}

func (g *EventSystem) Key(key engi.Key, modifier engi.Modifier, action engi.KeyAction) {
	if key == engi.Escape {
		engi.Exit()
	}
	event := &KeyEvent{key, modifier, action}
	for _, e := range g.keyEventChannels {
		e <- event
	}
}

func (g *EventSystem) Type(char rune) {
	event := &TypeKeyEvent{rune}
	for _, e := range g.typeKeyEventChannels {
		e <- event
	}
}
