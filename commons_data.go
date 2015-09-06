package lbricks

import (
	"github.com/ungerik/go3d/vec2"
	"github.com/Rubentxu/lbricks/engi"
)


type Transform struct {
	Position *vec2.T
	Scale  	*vec2.T
	Rotation float32
}

type View struct {
	*Transform
	Anchor   *vec2.T
	Color    uint32
	Alpha    float32
	Region   *engi.Region
}