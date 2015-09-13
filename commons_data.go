package lbricks



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
	Position() []float32
}

type Scalable interface {
	Scale() []float32
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

type Drawable interface {
	View() (float32, float32, float32, float32)
	Texture() uint
}

type Displayable interface {
	Visible() bool
}

type Resizable interface {
	Width() float32
	Height() float32
}

type Colorize interface {
	Color() (uint32, uint32, uint32)
	Alpha() float32
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
	Resizable
	Shape()  Shape
	Anchor() [2]float32
}

type Layout interface {
	Direction() Direction

}

type View interface {
	Transformer
	Drawable
	Element
	Layer() uint
}

type Batch interface {
	Begin()
	Draw(r Drawable, x, y, originX, originY, scaleX, scaleY, rotation float32, color uint32, transparency float32)
	End()
	SetProjection(width, height float32)
}