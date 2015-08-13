package lbricks

import (
	"github.com/Rubentxu/lbricks/engi"
)

type ClockEvent struct {
	time Clock
}

type MouseEvent struct {
	PosX, PosY   float32
	AmountScroll float32
	Action       engi.MouseAction
}

type KeyboardEvent struct {
	Key      engi.Key
	Modifier engi.Modifier
	Action   engi.KeyAction
}
