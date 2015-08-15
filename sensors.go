package lbricks

import (
	"github.com/Rubentxu/lbricks/engi"
	"github.com/trustmaster/goflow"
)

type ClockChannel interface {
	OnStep() chan<- *ClockEvent
}

type SimpleInChannel interface {
	OnIn() chan<- interface{} // The writeable end of the channel.
}

type SimpleOutChannel interface {
	OnOut() <-chan interface{} // The readable end of the channel.
}

type ISensor interface {
	SimpleInChannel
	SimpleOutChannel
	ClockChannel
}

type Sensor struct {
	name string
	Type string
	Step      <-chan *ClockEvent // input port
	frequency int
}

func (s *Sensor) OnStep(event *ClockEvent) {
	if()
}

type MouseSensor struct {
	flow.Component // component "superclass" embedded
	*Sensor
	In     <-chan *MouseEvent // input port
	Out    chan<- *MouseEvent // output port
	action engi.MouseAction
}

func NewMouseSensor(n string, f int,a engi.MouseAction) *MouseSensor {
	return &MouseSensor{name: n, Type: "MouseSensor" frequency : f, action: a}
}

func (ms *MouseSensor) OnIn(event *MouseEvent) {
	if event.Action == ms.Action {
		ms.Out <- event
	}
}

type KeyboardSensor struct {
	flow.Component                       // component "superclass" embedded
	In             <-chan *KeyboardEvent // input port
	Out            chan<- *KeyboardEvent // output port

	action  engi.KeyAction
	key     engi.Key
	allKeys bool
}

func NewKeyboardSensor() *KeyboardSensor {
	return &KeyboardSensor{allKeys: true}
}

func NewKeyboardSensor(a engi.KeyAction, k engi.Key) *KeyboardSensor {
	return &KeyboardSensor{action: a, key: k, allKeys: false}
}

func (ks *KeyboardSensor) OnIn(event *KeyboardEvent) {
	if ks.allKeys {
		ks.Out <- event
	} else if event.Action == ks.action && event.Key == ks.key {
		ks.Out <- event
	}
}
