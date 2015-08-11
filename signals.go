package lbricks

import (
	"github.com/ajhager/engi"
)

type MouseSignal struct {
	PosX, PosY float32
	Action     engi.Action
}

type KeyboardSignal struct {
	Key      engi.Key
	Modifier engi.Modifier
	Action   engi.Action
}
