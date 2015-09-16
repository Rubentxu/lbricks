package lbricks

import (
	"github.com/Rubentxu/lbricks/goflow"
	"github.com/ungerik/go3d/float64/vec2"
)

type MotionActuator struct {
	flow.Component
	velocity  		vec2.T;
	force 			vec2.T;
	impulse  		vec2.T;
	angularVelocity 	float32;
	torque 				float32;
	angularImpulse		float32;
	fixedRotation 		bool;
	limitVelocityX 		float32;
	limitVelocityY 		float32
	PulseState           <-chan bool // input port
	Res            chan<- string // output port
}