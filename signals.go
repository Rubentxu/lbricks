package lbricks

import (
	"github.com/Rubentxu/lbricks/engi"
)

type MouseSignal struct {
	PosX, PosY   float32
	AmountScroll float32
	Event        engi.MouseEvent
}

type KeyboardSignal struct {
	Key      engi.Key
	Modifier engi.Modifier
	Action   engi.Action
}
