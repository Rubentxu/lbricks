package lbricks

import (
	"github.com/Rubentxu/lbricks/engi"
)

type EventPacket struct {
	event *Event
	Code  int
	Data  interface{}
}

type Event struct {
	eventType string
}

func (e *Event) Type() string {
	return e.eventType
}

type StepEvent struct {
	*Event
	step    float64
	numStep int
}

type MouseEvent struct {
	*Event
	PosX, PosY   float32
	AmountScroll float32
	Action       engi.MouseAction
}

type KeyboardEvent struct {
	*Event
	Key      engi.Key
	Modifier engi.Modifier
	Action   engi.KeyAction
}
