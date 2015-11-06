package lbricks

import "github.com/ungerik/go3d/vec2"

type Typable interface  {
	Type() .(type)

}
type Poolable  interface {
	Reset()
}

type Disposable interface {
	Dispose()
}

type Renderable interface {
	Render(Batch)
}

type Positionable interface {
	Position() *vec2.T
}

type Scalable interface {
	Scale() *vec2.T
}

type Rotable interface {
	Rotation() float32
}

type Identificable interface {
	Id() uint
}

type Nombrable interface {
	Name() string
}

type Tagger interface {
	Tags() []string
}

type Displayable interface {
	Visible() bool
}

type Colorize interface {
	RGB() uint32
	Alpha() float32
}

type Drawable interface {
	Texture() uint
	Width() float32
	Height() float32
	View() (float32, float32, float32, float32)
}

type Recognizable interface {
	Identificable
	Nombrable
	Tagger
}

type Transformer interface {
	Positionable
	Scalable
	Rotable
}

type Element interface {
	Colorize
	Drawable
	Shape()  Shape
	Anchor() [2]float32
}

type Layout interface {
	Direction() Direction
}

type View interface {
	Transformer
	Element
	Layer() uint
}

type Batch interface {
	Begin()
	Draw(r Drawable, x, y, originX, originY, scaleX, scaleY, rotation float32, color uint32, transparency float32)
	End()
	SetProjection(width, height float32)
}

type Exception string

type Event interface {
	Identificable
	Nombrable
}

// Signal
type ISignal interface {
	Subscribe(f func(e Event)) chan Event
	OnNext(event Event)
	OnError(error Exception)
	OnCompleted()
}


