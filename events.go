package lbricks

import (
	"github.com/Rubentxu/lbricks/engi"
)

type EventPacket struct {
	Code  int
	event interface{}
}

type EventType string

type StepEvent struct {
	eventType EventType
	step      float64
	numStep   int
}

type MouseEvent struct {
	eventType    EventType
	PosX, PosY   float32
	AmountScroll float32
	Action       engi.MouseAction
}

type KeyboardEvent struct {
	eventType EventType
	Key       engi.Key
	Modifier  engi.Modifier
	Action    engi.KeyAction
}
