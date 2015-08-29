package lbricks

import (
	"github.com/Rubentxu/lbricks/engi"
)

type EventPacket struct {
	Code int
	Data interface{}
}

// StepEvent
type StepEvent struct {
	step    float64
	numStep int
}

// MouseEvent
type MouseEvent struct {
	PosX, PosY   float32
	AmountScroll float32
	Action       engi.MouseAction
}

// KeyboardEvent
type KeyboardEvent struct {
	Key      engi.Key
	Modifier engi.Modifier
	Action   engi.KeyAction
}