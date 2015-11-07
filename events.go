package lbricks

type EventPacket struct {
	Event interface{}

}

type PreloadEvent struct {}
type SetupEvent struct {}
type CloseEvent struct {}
type UpdateEvent struct {
	DeltaTime float32
}
type RenderEvent struct {}
type ResizeEvent struct {
	Width, Height int
}

// StepEvent
type StepEvent struct {
	Step    float64
	NumStep uint32
}

// MouseEvent
type MouseEvent struct {
	PosX, PosY float32
	Action     MouseAction
}

type ScrollEvent struct {
	AmountScroll float32
	Action     MouseAction
}

// KeyboardEvent
type KeyEvent struct {
	Key      Key
	Modifier Modifier
	Action   KeyAction
}

type TypeKeyEvent struct {
	Char rune
}

