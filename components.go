package lbricks

import (
	"github.com/ungerik/go3d/vec2"
	"github.com/Rubentxu/lbricks/goflow"
	"fmt"
)

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
	Action           <-chan string // input port
	position *vec2.T
	scale    *vec2.T
	rotation float32
}

func (t *Transform) Position() *vec2.T {
	return t.position
}

func (t *Transform) Scale() *vec2.T {
	return t.scale
}

func (t *Transform) Rotation() float32 {
	return t.rotation
}

func (t *Transform) OnAction(action string) {
	fmt.Printf("Transform action, %s!", action)
}


// Color Component
type Color struct {
	rgb uint32
	alpha   float32
}

func (c *Color) RGB() uint32 {
	return c.rgb
}

func (c *Color) Alpha() float32 {
	return c.alpha
}


type Shape []float32

var (
	Rect = 		Shape{0.0,0.0}
	Oval = 		Shape{0,0}
	Square = 	Shape{0}
	Circle = 	Shape{0}
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
	Shape
	view Drawable
	Anchor 	[2]float32
	layer   uint
	width, height float32
}

func (v *Sprite) Render(batch Batch) {
	batch.Draw(v.view, v.Position()[0], v.Position()[1], v.Anchor[0], v.Anchor[1], v.Scale()[0],
		v.Scale()[1], v.Rotation(), v.RGB(), v.Alpha())
}

func (s *Sprite) Width() float32 {
	return s.width
}

func (s *Sprite) Height() float32 {
	return s.height
}

