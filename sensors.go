package lbricks

import (
	"github.com/Rubentxu/lbricks/goflow"
)


/* SensorBase */
type Sensor struct {
	flow.Component
	Name      string
	EventType string
	StepEvent      <-chan StepEvent  // input port
	Out       chan <- interface{} // output port
	Frequency int
}

func (s *Sensor) AddToGraph(graph flow.Graph) {
	graph.Add(s,s.Name)
	var stepEvent string = "StepEvent"
	stepEventIn := make(chan *StepEvent)
	graph.MapInPort(s.Name+stepEvent,s.Name,stepEvent)
	graph.SetInPort(s.Name+stepEvent,stepEventIn)
	graph.MapInPort(s.Name+s.EventType,s.Name,s.EventType)
}

/* MouseSensor */
type MouseSensor struct {
	Sensor
	MouseEvent <-chan *MouseEvent
	action MouseAction
}

func NewMouseSensor(name string, freq int, action MouseAction) *MouseSensor {
	sensor := new(MouseSensor)
	sensor.Name = name
	sensor.Frequency = freq
	sensor.action = action
	return sensor
}

func (ms *MouseSensor) AddToGraph(graph flow.Graph) {
	ms.Sensor.AddToGraph(graph)
	mouseEventIn := make(chan *MouseEvent)
	graph.SetInPort(ms.Name+ms.EventType,mouseEventIn)
}

func (ms *MouseSensor) OnMouseEvent(event *MouseEvent) {
	if event.Action == ms.action {
		ms.Out <- event
	}
}

// KeyboardSensor
type KeyboardSensor struct {
	Sensor                        // component "superclass" embedded
	KeyboardEvent      <-chan *KeyEvent  // input port
	Out     chan <- *KeyEvent // output port
	action  KeyAction
	key     Key
	allKeys bool
}

func NewKeyboardSensor(a KeyAction, k Key) *KeyboardSensor {
	return &KeyboardSensor{action: a, key: k, allKeys: false}
}

func (ks *KeyboardSensor) OnKeyboardEvent(event *KeyEvent) {
	if ks.allKeys {
		ks.Out <- event
	} else if event.Action == ks.action && event.Key == ks.key {
		ks.Out <- event
	}
}
