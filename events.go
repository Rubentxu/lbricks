package lbricks

import (
	"github.com/Rubentxu/lbricks/engi"
)

type Event interface {
	Type
}

type EventBase struct {
	eventType string
}

func (e *EventBase) Type() string {
	return e.eventType
}

type StepEvent struct {
	*EventBase
	step    float64
	numStep int
}

type MouseEvent struct {
	*EventBase
	PosX, PosY   float32
	AmountScroll float32
	Action       engi.MouseAction
}

type KeyboardEvent struct {
	*EventBase
	Key      engi.Key
	Modifier engi.Modifier
	Action   engi.KeyAction
}
