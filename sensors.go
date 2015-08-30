package lbricks

import (
	"github.com/Rubentxu/lbricks/engi"
	"github.com/trustmaster/goflow"
)

type EventType string

type Sensor interface {
	SimpleInChannel
	SimpleOutChannel
	ClockChannel
}

/* SensorBase */
type sensor struct {
	flow.Component
	name      string
	eventType EventType
	Step      <-chan EventPacket  // input port
	In        <-chan EventPacket  // input port
	Out       chan <- EventPacket // output port
	frequency int
}

func (s *sensor) Name() string {
	return s.name
}

func (s *sensor) EventType() EventType {
	return s.eventType
}

func (s *sensor) Frequency() int {
	return s.frequency
}

func (s *sensor) OnCompleteStep() chan <- *StepEvent {
	return nil
}

func (s *sensor) AddSensor(graph flow.Graph, inputs map[EventType] []chan *EventPacket) {
	graph.Add(s, s.Name())
	inputs[]

}

/* MouseSensor */
type MouseSensor struct {
	*sensor
	action engi.MouseAction
}

func NewMouseSensor(n string, f int, a engi.MouseAction) *MouseSensor {
	sensor := new(MouseSensor)
	sensor.name = n
	sensor.eventType = "MouseEvent"
	sensor.frequency = f
	return sensor
}

func (ms *MouseSensor) OnIn(event *MouseEvent) {
	if event.Action == ms.action {
		ms.Out <- event
	}
}

// KeyboardSensor
type KeyboardSensor struct {
	*sensor                        // component "superclass" embedded
	In      <-chan *KeyboardEvent  // input port
	Out     chan <- *KeyboardEvent // output port
	action  engi.KeyAction
	key     engi.Key
	allKeys bool
}

func NewKeyboardSensor() *KeyboardSensor {
	return &KeyboardSensor{allKeys: true}
}

func NewKeyboardSensor2(a engi.KeyAction, k engi.Key) *KeyboardSensor {
	return &KeyboardSensor{action: a, key: k, allKeys: false}
}

func (ks *KeyboardSensor) OnIn(event *KeyboardEvent) {
	if ks.allKeys {
		ks.Out <- event
	} else if event.Action == ks.action && event.Key == ks.key {
		ks.Out <- event
	}
}
