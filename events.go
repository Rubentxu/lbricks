package lbricks

import (
	"github.com/Rubentxu/lbricks/engi"
)

type EventPacket struct {
	Code int
	Data *Event
}

type Event interface {
	Type
}

type StepEvent struct {
	step    float64
	numStep int
}

func (e *StepEvent) Type() {
	return "StepEvent"
}

type MouseEvent struct {
	PosX, PosY   float32
	AmountScroll float32
	Action       engi.MouseAction
}

func (e *MouseEvent) Type() {
	return "MouseEvent"
}

type KeyboardEvent struct {
	Key      engi.Key
	Modifier engi.Modifier
	Action   engi.KeyAction
}

func (e *KeyboardEvent) Type() {
	return "KeyboardEvent"
}
