package EngiFlow

import (
	"github.com/ajhager/engi"
	"github.com/trustmaster/goflow"
)

type MouseSensor struct {
	flow.Component                     // component "superclass" embedded
	In             <-chan *MouseSignal // input port
	Out            chan<- *MouseSignal // output port
}

func (g *MouseSensor) OnIn(ms *MouseSignal) {
	if ms.Action == engi.PRESS {
		g.Out <- ms
	}
}
