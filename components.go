package lbricks

import "github.com/ungerik/go3d/vec2"

// Identity Component
type Identity struct {
	id   uint
	name string
	tags []string
}

func (i *Identity) Id() uint {
	return i.id
}

func (i *Identity) Name() string {
	return i.name
}

func (i *Identity) Tags() []string {
	return i.tags
}


// Transform Component
type Transform struct {
	position *vec2.T
	scale    *vec2.T
	rotation float32
}

func (t *Transform) Position() []float32 {
	return t.position().([2]float64)
}

func (t *Transform) Scale() []float32 {
	return t.scale().([2]float64)
}

func (t *Transform) Rotation() float32 {
	return t.rotation
}


// Color Component
type Color struct {
	r, g, b uint32
	alpha   float32
}

func (c *Color) Color() (uint32, uint32, uint32) {
	return c.r, c.g, c.b
}

func (c *Color) Alpha() float32 {
	return c.alpha
}


// Superficie Component
type Superficie struct {
	width, height float32
}

func (a *Superficie) Width() float32 {
	return a.width
}

func (a *Superficie) Height() float32 {
	return a.height
}

type Shape []float32

const (
	Rect = 		[2]Shape{0,0}
	Oval = 		[2]Shape{0,0}
	Square = 	[1]Shape{0}
	Circle = 	[1]Shape{0}
)

type Direction uint

const (
	Up = Direction(0)
	Down = Direction(1)
	Left = Direction(2)
	Right = Direction(3)
	Inward = Direction(4)
	Outward = Direction(5)
)

type Sprite struct {
	Identity
	Transform
	Color
	Superficie
	Shape
	view Drawable
	Anchor 	[2]float32
}

func (v *Sprite) Render(batch Batch) {
	batch.Draw(v, v.Position()[0], v.Position()[1], v.Anchor[0], v.Anchor[1], v.Scale()[0],
		v.Scale()[1], v.Rotation(), v.Color, v.Alpha)
}
