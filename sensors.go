package lbricks

import (
	"github.com/Rubentxu/lbricks/engi"
	"github.com/trustmaster/goflow"
)

type TickInChannel interface {
	OnTick() chan<- *ClockSignal
}

type SimpleInChannel interface {
	OnIn() chan<- interface{} // The writeable end of the channel.
	Close()                   // Closes the channel. It is an error to write to In() after calling Close().
}

type SimpleOutChannel interface {
	OnOut() <-chan interface{} // The readable end of the channel.
}

type Sensor interface {
	SimpleInChannel
	SimpleOutChannel
	TickInChannel
}

type MouseSensor struct {
	flow.Component                     // component "superclass" embedded
	In             <-chan *MouseSignal // input port
	Out            chan<- *MouseSignal // output port

	Event engi.MouseEvent
}

func (ms *MouseSensor) OnIn(signal *MouseSignal) {
	if signal.Event == ms.Event {
		ms.Out <- signal
	}
}

type KeyboardSensor struct {
	flow.Component                        // component "superclass" embedded
	In             <-chan *KeyboardSignal // input port
	Out            chan<- *KeyboardSignal // output port

	action  engi.Action
	keyCode engi.Key
	allKeys bool
}

func (ks *KeyboardSensor) OnIn(signal *KeyboardSignal) {
	if ks.allKeys {
		ks.Out <- signal
	} else if signal.Action == ks.action && signal.Key == ks.keyCode {
		ks.Out <- signal
	}
}
