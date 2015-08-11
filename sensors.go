package lbricks

import (
	"github.com/ajhager/engi"
	"github.com/trustmaster/goflow"
)

type MouseSensor struct {
	flow.Component                     // component "superclass" embedded
	In             <-chan *MouseSignal // input port
	Out            chan<- *MouseSignal // output port

	Action engi.Action
}

func (ms *MouseSensor) OnIn(signal *MouseSignal) {
	if signal.Action == ms.Action {
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

func (ks *KeyboardSensor) NewKeyboardSensor(action engi.Action, keyCode engi.Key) {
	return
}

func (ks *KeyboardSensor) OnIn(signal *KeyboardSignal) {
	if ks.allKeys {
		ks.Out <- signal
	} else if signal.Action == ks.action && signal.Key == ks.keyCode {
		ks.Out <- signal
	}
}
