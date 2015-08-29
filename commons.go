package lbricks

type Type interface {
	Type() string
}

type ClockChannel interface {
	OnCompleteStep() chan <- *StepEvent
}

type SimpleInChannel interface {
	OnIn(event EventPacket) chan <- interface{} // The writeable end of the channel.
}

type SimpleOutChannel interface {
	OnOut(event EventPacket) <-chan interface{} // The readable end of the channel.
}
