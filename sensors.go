package lbricks

import (
	"github.com/Rubentxu/lbricks/engi"
	"github.com/trustmaster/goflow"
)

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
