package lbricks

import (
	"github.com/Rubentxu/lbricks/engi"
)

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
	AmountScroll loat32
	Action       engi.MouseAction
}

type KeyboardEvent struct {
	*Event
	Key      engi.Key
	Modifier engi.Modifier
	Action   engi.KeyAction
}
