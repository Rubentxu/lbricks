package lbricks

import (
	"github.com/Rubentxu/lbricks/engi"
	"github.com/Rubentxu/lbricks/goflow"
)

type EventType string


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
	action engi.MouseAction
}

func NewMouseSensor(name string, freq int, action engi.MouseAction) *MouseSensor {
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
	KeyboardEvent      <-chan *KeyboardEvent  // input port
	Out     chan <- *KeyboardEvent // output port
	action  engi.KeyAction
	key     engi.Key
	allKeys bool
}

func NewKeyboardSensor(a engi.KeyAction, k engi.Key) *KeyboardSensor {
	return &KeyboardSensor{action: a, key: k, allKeys: false}
}

func (ks *KeyboardSensor) OnKeyboardEvent(event *KeyboardEvent) {
	if ks.allKeys {
		ks.Out <- event
	} else if event.Action == ks.action && event.Key == ks.key {
		ks.Out <- event
	}
}
