package lbricks

import (
	"github.com/Rubentxu/lbricks/engi"
	"github.com/trustmaster/goflow"
)

type Type interface {
	Type() string
}

type ClockChannel interface {
	OnCompleteStep() chan<- *StepEvent
}

type SimpleInChannel interface {
	OnIn() chan<- interface{} // The writeable end of the channel.
}

type SimpleOutChannel interface {
	OnOut() <-chan interface{} // The readable end of the channel.
}

type Sensor interface {
	Type
	SimpleInChannel
	SimpleOutChannel
	ClockChannel
	EventType() string
}

type SensorBase struct {
	flow.Component
	name       string
	sensorType string
	eventType  string
	Step       <-chan Event // input port
	In         <-chan Event // input port
	Out        chan<- Event // output port
	frequency  int
}

func (s *SensorBase) Name() string {
	return s.name
}

func (s *SensorBase) Type() string {
	return s.sensorType
}

func (s *SensorBase) EventType() string {
	return s.eventType
}

func (s *SensorBase) Frequency() int {
	return s.frequency
}

func (s *SensorBase) OnCompleteStep() chan<- *StepEvent {
	return nil
}

type MouseSensor struct {
	*SensorBase
	action engi.MouseAction
}

func NewMouseSensor(n string, f int, a engi.MouseAction) *MouseSensor {
	sensor := new(MouseSensor)
	sensor.name = n
	sensor.sensorType = "MouseSensor"
	return sensor
}

func (ms *MouseSensor) OnIn(event *MouseEvent) {
	if event.Action == ms.action {
		ms.Out <- event
	}
}

type KeyboardSensor struct {
	*SensorBase                       // component "superclass" embedded
	In          <-chan *KeyboardEvent // input port
	Out         chan<- *KeyboardEvent // output port

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
