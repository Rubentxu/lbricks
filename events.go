package lbricks

import (
	"github.com/Rubentxu/lbricks/engi"
)

type PreloadEvent struct {}
type SetupEvent struct {}
type CloseEvent struct {}
type UpdateEvent struct {
	dt float32
}
type RenderEvent struct {}
type ResizeEvent struct {
	Width, Height int
}

// StepEvent
type StepEvent struct {
	Step    float64
	NumStep int
}

// MouseEvent
type MouseEvent struct {
	PosX, PosY   float32
	Action       engi.MouseAction
}

type ScrollEvent struct {
	AmountScroll float32
}

// KeyboardEvent
type KeyEvent struct {
	Key      engi.Key
	Modifier engi.Modifier
	Action   engi.KeyAction
}

type TypeKeyEvent struct {
	Char rune
}