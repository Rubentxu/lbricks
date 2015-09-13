package lbricks

type Poolable  interface {
	Reset()
}

type Disposable interface{
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
	Id() 	uint
}

type Nombrable interface {
	Name()	string
}

type Tagger interface {
	Tags()	[]string
}

type Texturizer interface {
	Texture()	int
}

type Drawable interface {
	View() (float32, float32, float32, float32)
}

type Resizable interface {
	Width() float32
	Height() float32
}

type Colour interface {
	Color() uint32
	Alpha()	float32
}

type Batch interface {
	Begin()
	Draw(r Drawable, x, y, originX, originY, scaleX, scaleY, rotation float32, color uint32, transparency float32)
	End()
	SetProjection(width, height float32)
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

type View interface {
	Renderable
	Transformer
	Drawable
	Recognizable
	Anchor() []float32
	Layer() uint
}

type Sprite struct {
	Transform Transformer
	id 		uint
	Name 	string
	Anchor   []float32
	Color    uint32
	Alpha    float32
	Region   Drawable
	Layer   uint
}

func (v *Sprite) Render(batch Batch) {
	batch.Draw(v.Region, v.Transform.Position()[0], v.Transform.Position()[1], v.Anchor[0], v.Anchor[1], v.Transform.Scale()[0],
		v.Transform.Scale()[1], v.Transform.Rotation(), v.Color, v.Alpha)
}