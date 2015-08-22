package lbricks

type broadcast struct {
	c chan broadcast
	v Event
}

type Broadcaster struct {
	ListenMouseEvent    chan chan (chan broadcast)
	ListenKeyboardEvent chan chan (chan broadcast)
	Send                chan<- Event
}

// create a new broadcaster object.
func NewBroadcaster() Broadcaster {
	listenMouseEvent := make(chan (chan (chan broadcast)))
	listenKeyboardEvent := make(chan (chan (chan broadcast)))
	send := make(chan Event)
	go func() {
		currc := make(chan broadcast, 1)
		for {
			select {
			case v := <-send:
				c := make(chan broadcast, 1)
				b := broadcast{c: c, v: v}
				currc <- b
				currc = c
			case m := <-listenMouseEvent:
				m <- currc
			case k := <-listenKeyboardEvent:
				k <- currc
			}
		}
	}()
	return Broadcaster{
		ListenMouseEvent:    listenMouseEvent,
		ListenKeyboardEvent: listenKeyboardEvent,
		Send:                send,
	}
}

// start listening to the broadcasts.
func (b Broadcaster) Register(sensor Sensor) {
	c := make(chan chan broadcast, 0)

	switch sensor.Type() {
	case "MouseSensor":
		b.ListenMouseEvent <- c
	case "KeyboardSensor":
		b.ListenKeyboardEvent <- c
	}
	rv := Receiver{<-c, sensor}
	go rv.listen()
}

// broadcast a value to all listeners.
func (b Broadcaster) Write(v Event) { b.Send <- v }

type Receiver struct {
	C      chan broadcast
	sensor Sensor
}

func (r *Receiver) Read() Event {
	b := <-r.C
	v := b.v
	r.C <- b
	r.C = b.c
	return v
}

func (r *Receiver) listen() {
	for v := r.Read(); r.sensor.EventType() == v.Type(); v = r.Read() {
		r.sensor.OnIn(v)
	}
}
