package lbricks

type Disposable interface{
	Dispose()
}

type Drawable interface {
	Texture() int
	Width() float32
	Height() float32
	View() (float32, float32, float32, float32)
}

type Batch interface {
	Begin()
	Draw(r Drawable, x, y, originX, originY, scaleX, scaleY, rotation float32, color uint32, transparency float32)
	End()
	SetProjection(width, height float32)
}

type Renderer interface {
	Render(Batch)
}

type Transform interface {
	Position() []float32
	Scale()  	[]float32
	Rotation() float32
}


type View interface {
	Renderer
	Transform
	Drawable
	Id() uint
	Name() string
	Anchor() []float32
	Color() uint32
	Alpha()	float32
	Layer() uint

}

type Sprite struct {
	Transform Transform
	id 		uint
	Name 	string
	Anchor   []float32
	Color    uint32
	Alpha    float32
	Region   Drawable
	Layer   uint
}

func (v *Sprite) Render(batch Batch) {
	batch.Draw(v.Region, v.Transform.Position[0], v.Transform.Position[1], v.Anchor[0], v.Anchor[1], v.Transform.Scale[0],
		v.Transform.Scale[1], v.Transform.Rotation, v.Color, v.Alpha)
}