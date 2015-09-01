package lbricks

import (
	"github.com/Rubentxu/lbricks/goflow"
	"reflect"
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
	switch input.Channel.Type().Name() {
	case "PreloadEvent" :
		g.preloadEventChannnels = append(g.preloadEventChannnels, input.Channel.Interface().(chan *PreloadEvent))
	case "SetupEvent" :
		g.setupEventChannels = append(g.setupEventChannels, input.Channel.Interface().(chan *SetupEvent))
	case "CloseEvent" :
		g.closeEventChannels = append(g.closeEventChannels, input.Channel.Interface().(chan *CloseEvent))
	case "UpdateEvent" :
		g.updateEventChannels = append(g.updateEventChannels, input.Channel.Interface().(chan *UpdateEvent))
	case "RenderEvent" :
		g.renderEventChannels = append(g.renderEventChannels, input.Channel.Interface().(chan *RenderEvent))
	case "ResizeEvent" :
		g.resizeEventChannels = append(g.resizeEventChannels, input.Channel.Interface().(chan *ResizeEvent))
	case "StepEvent" :
		g.stepEventChannels = append(g.stepEventChannels, input.Channel.Interface().(chan *StepEvent))
	case "MouseEvent" :
		g.mouseEventChannels = append(g.mouseEventChannels, input.Channel.Interface().(chan *MouseEvent))
	case "ScrollEvent" :
		g.scrollEventChannels = append(g.scrollEventChannels, input.Channel.Interface().(chan *ScrollEvent))
	case "KeyEvent" :
		g.keyEventChannels = append(g.keyEventChannels, input.Channel.Interface().(chan *KeyEvent))
	case "TypeKeyEvent" :
		g.typeKeyEventChannels = append(g.typeKeyEventChannels, input.Channel.Interface().(chan *TypeKeyEvent))
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

func (g *EventSystem) Mouse(x, y float32, action interface{}) {
	a:= MouseAction(reflect.ValueOf(action).Interface().(int))
	event := &MouseEvent{x, y, a}
	for _, e := range g.mouseEventChannels {
		e <- event
	}
}

func (g *EventSystem) Scroll(amount float32) {
	var action MouseAction = WHEEL_UP
	if amount < 0 {
		action = WHEEL_DOWN
	}
	event := &ScrollEvent{amount, action}
	for _, e := range g.scrollEventChannels {
		e <- event
	}
}

func (g *EventSystem) Key(key interface{}, modifier interface{}, action interface{}) {
	if key == Escape {
		//Exit()
	}
	k := Key(reflect.ValueOf(key).Interface().(int))
	m := Modifier(reflect.ValueOf(modifier).Interface().(int))
	a := KeyAction(reflect.ValueOf(action).Interface().(int))
	event := &KeyEvent{k, m, a}
	for _, e := range g.keyEventChannels {
		e <- event
	}
}

func (g *EventSystem) Type(char rune) {
	event := &TypeKeyEvent{char}
	for _, e := range g.typeKeyEventChannels {
		e <- event
	}
}
