package lbricks

import (
	"github.com/Rubentxu/lbricks/engi"
	"github.com/trustmaster/goflow"
)

type Sensor interface {
	SimpleInChannel
	SimpleOutChannel
	ClockChannel
	EventType() string
}

/* SensorBase */
type SensorBase struct {
	flow.Component
	name      string
	eventType string
	Step      <-chan EventPacked // input port
	In        <-chan EventPacked // input port
	Out       chan<- EventPacked // output port
	frequency int
}

func (s *SensorBase) Name() string {
	return s.name
}

func (s *SensorBase) EventType() string {
	return s.eventType
}

func (s *SensorBase) Frequency() int {
	return s.frequency
}

func (s *SensorBase) OnCompleteStep() chan<- *StepEvent {
	return nil
}

/* MouseSensor */
type MouseSensor struct {
	*SensorBase
	action engi.MouseAction
}

func NewMouseSensor(n string, f int, a engi.MouseAction) *MouseSensor {
	sensor := new(MouseSensor)
	sensor.name = n
	sensor.sensorType = "MouseSensor"
	sensor.eventType = "MouseEvent"
	sensor.frequency = f
	return sensor
}

func (ms *MouseSensor) OnIn(event *MouseEvent) {
	if event.Action == ms.action {
		ms.Out <- event
	}
}

// KeyboardSensor
type KeyboardSensor struct {
	*SensorBase                       // component "superclass" embedded
	In          <-chan *KeyboardEvent // input port
	Out         chan<- *KeyboardEvent // output port
	action      engi.KeyAction
	key         engi.Key
	allKeys     bool
}

func NewKeyboardSensor() *KeyboardSensor {
	return &KeyboardSensor{allKeys: true}
}

func NewKeyboardSensor2(a engi.KeyAction, k engi.Key) *KeyboardSensor {
	return &KeyboardSensor{action: a, key: k, allKeys: false}
}

func (ks *KeyboardSensor) OnIn(event *KeyboardEvent) {
	if ks.allKeys {
		ks.Out <- event
	} else if event.Action == ks.action && event.Key == ks.key {
		ks.Out <- event
	}
}
