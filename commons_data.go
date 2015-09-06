package lbricks

import (
	"github.com/ungerik/go3d/vec2"
	"github.com/Rubentxu/lbricks/engi"
)

type Renderer interface {
	Render(engi.Batch)
}

type Transform struct {
	Position *vec2.T
	Scale  	*vec2.T
	Rotation float32
}

func NewTransform() *Transform{
	t:= new(Transform)
	t.Position = vec2.Zero
	t.Scale = vec2.UnitXY
	return t
}

type View struct {
	*Transform
	id 		uint
	Name 	string
	Anchor  *vec2.T
	Color   uint32
	Alpha   float32
	Layer   uint
	Region  *engi.Region
}

func NewView(id uint,name string)  {
	v:= new(View)
	v.id = id
	v.Name = name
	v.Anchor = vec2.Zero
	v.Transform = NewTransform()
}

func (v *View) Render(batch *engi.Batch) {
	batch.Draw(v.Region, v.Position[0], v.Position[1], v.Anchor[0], v.Anchor[1], v.Scale[0], v.Scale[1], v.Rotation, v.Color, v.Alpha)
}