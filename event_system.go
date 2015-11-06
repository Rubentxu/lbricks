package lbricks

import (
	"github.com/Rubentxu/lbricks/goflow"
	"reflect"
	"github.com/tcard/functional"
)


type EventSystem struct {
	PreloadEventChannnels 	[]chan *PreloadEvent
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

	EventChannels 			map[string] [] *func(Event) Event
}



func CreateEventSystem(capacity int) *EventSystem  {
	eventSystem := new(EventSystem)
	eventSystem.PreloadEventChannnels = make([]chan *PreloadEvent, 0, capacity)
	eventSystem.setupEventChannels =	make([]chan *SetupEvent, 0, capacity)
	eventSystem.closeEventChannels = 	make([]chan *CloseEvent, 0, capacity)
	eventSystem.updateEventChannels = 	make([]chan *UpdateEvent, 0, capacity)
	eventSystem.renderEventChannels = 	make([]chan *RenderEvent, 0, capacity)
	eventSystem.resizeEventChannels = 	make([]chan *ResizeEvent, 0, capacity)
	eventSystem.stepEventChannels = 	make([]chan *StepEvent, 0, capacity)
	eventSystem.mouseEventChannels = 	make([]chan *MouseEvent, 0, capacity)
	eventSystem.scrollEventChannels = 	make([]chan *ScrollEvent, 0, capacity)
	eventSystem.keyEventChannels = 		make([]chan *KeyEvent, 0, capacity)
	eventSystem.typeKeyEventChannels = 	make([]chan *TypeKeyEvent, 0, capacity)
	return eventSystem
}

func (g *EventSystem) RegisterInputChannel(input flow.Port) {
	event := input.Channel.Interface()
	switch event.(type) {
	case chan *PreloadEvent :
		g.PreloadEventChannnels = append(g.PreloadEventChannnels, input.Channel.Interface().(chan *PreloadEvent))
	case chan *SetupEvent :
		g.setupEventChannels = append(g.setupEventChannels, input.Channel.Interface().(chan *SetupEvent))
	case chan *CloseEvent :
		g.closeEventChannels = append(g.closeEventChannels, input.Channel.Interface().(chan *CloseEvent))
	case chan *UpdateEvent :
		g.updateEventChannels = append(g.updateEventChannels, input.Channel.Interface().(chan *UpdateEvent))
	case chan *RenderEvent :
		g.renderEventChannels = append(g.renderEventChannels, input.Channel.Interface().(chan *RenderEvent))
	case chan *ResizeEvent :
		g.resizeEventChannels = append(g.resizeEventChannels, input.Channel.Interface().(chan *ResizeEvent))
	case chan *StepEvent :
		g.stepEventChannels = append(g.stepEventChannels, input.Channel.Interface().(chan *StepEvent))
	case chan *MouseEvent :
		g.mouseEventChannels = append(g.mouseEventChannels, input.Channel.Interface().(chan *MouseEvent))
	case chan *ScrollEvent :
		g.scrollEventChannels = append(g.scrollEventChannels, input.Channel.Interface().(chan *ScrollEvent))
	case chan *KeyEvent :
		g.keyEventChannels = append(g.keyEventChannels, input.Channel.Interface().(chan *KeyEvent))
	case chan *TypeKeyEvent :
		g.typeKeyEventChannels = append(g.typeKeyEventChannels, input.Channel.Interface().(chan *TypeKeyEvent))
	}
}

func (es *EventSystem) fromEvent(eventType string ) func( func(event Event) Event) {

	return func(fn func(Event) Event) {
		es.addEventListener(eventType, func(event Event) Event {
			return fn(event)
		})
	}
}


func (es *EventSystem) addEventListener(eventType string, fn func(Event) Event)  {
	if _,ok := es.EventChannels[eventType]; !ok {
		es.EventChannels[eventType] =  make([]func(Event) Event)
	}
	es.EventChannels[eventType] = append(es.EventChannels[eventType],fn)
}


type Stream struct {
	fns   [] func(Event) Event
}


func (s *Stream) Map (fn func (Event) Event) *Stream  {
		fn()
}



func  RegisterPort(port chan Event) {
	g := EventSystem {}

	switch port.(type) {
	case chan PreloadEvent :
		g.PreloadEventChannnels = append(g.PreloadEventChannnels, port.(PreloadEvent))

	}
}

func (g *EventSystem) Preload() {
	event := &PreloadEvent{}
	for _, e := range g.PreloadEventChannnels {
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
