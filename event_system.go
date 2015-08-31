package lbricks

import (
	"github.com/Rubentxu/lbricks/engi"
	"github.com/Rubentxu/lbricks/goflow"
)


type EventSystem struct {
	preloadChannnels	 	[]chan *PreloadEvent
	setupEventChannels	 	[]chan *SetupEvent
	closeEventChannels	 	[]chan *CloseEvent
	updateEventChannels	 	[]chan *UpdateEvent
	renderEventChannels	 	[]chan *RenderEvent
	resizeEventChannels	 	[]chan *ResizeEvent
	stepChannels 			[]chan *StepEvent
	mouseChannels 			[]chan *MouseEvent
	scrollEventChannels	 	[]chan *ScrollEvent
	keyChannels 			[]chan *KeyEvent
	typeKeyEventChannels 	[]chan *TypeKeyEvent
}

func (g *EventSystem) RegisterInputChannel(input flow.Port) {
	switch input.Port {
	case "PreloadEvent" :
		append(g.preloadChannnels, input.Channel)
	case "SetupEvent" :
		append(g.setupEventChannels, input.Channel)
	case "CloseEvent" :
		append(g.closeEventChannels, input.Channel)
	case "UpdateEvent" :
		append(g.updateEventChannels, input.Channel)
	case "RenderEvent" :
		append(g.renderEventChannels, input.Channel)
	case "ResizeEvent" :
		append(g.resizeEventChannels, input.Channel)
	case "StepEvent" :
		append(g.stepChannels, input.Channel)
	case "MouseEvent" :
		append(g.mouseChannels, input.Channel)
		append(g.scrollEventChannels, input.Channel)
	case "KeyboardEvent" :
		append(g.keyChannels, input.Channel)
		append(g.typeKeyEventChannels, input.Channel)
	}
}

func (g *EventSystem) Preload() {
	event := &PreloadEvent{}
	for _, e := range g.preloadChannnels {
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
	for _, e := range g.stepChannels {
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
	for _, e := range g.mouseChannels {
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
	for _, e := range g.keyChannels {
		e <- event
	}
}

func (g *EventSystem) Type(char rune) {
	event := &TypeKeyEvent{rune}
	for _, e := range g.typeKeyEventChannels {
		e <- event
	}
}
