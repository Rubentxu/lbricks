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
	base *event
	Key      Key
	Modifier Modifier
	Action   KeyAction
}

type event struct {
	id uint
	name string
}

type TypeKeyEvent struct {
	Char rune
}

func (ke *KeyEvent) ID()  uint {
	return ke.base.id
}


func (ke *KeyEvent) Name() string {
	return ke.base.name
}


type Signal struct {
	port chan Event
}

func (s *Signal) OnNext(event Event)  {
	s.port <- event
}

func (s *Signal) Subscribe(f func(e Event)) chan Event {
	if s.port == nil { s.port = make(chan Event) }
	go func() {
		defer s.Dispose()
		for event := range s.port {
			f(event)
		}
	}()
	return s.port
}

func (s *Signal) Dispose()  {
	close(s.port)
}
