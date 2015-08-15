package lbricks

import (
	"fmt"
	"time"
)

type broadcast struct {
	c chan broadcast
	v interface{}
}

type Broadcaster struct {
	// private fields:
	ListenInt    chan chan (chan broadcast)
	ListenString chan chan (chan broadcast)
	Sendc        chan<- interface{}
}

// create a new broadcaster object.
func NewBroadcaster() Broadcaster {
	listenMouse := make(chan (chan (chan broadcast)))
	listenKeyboard := make(chan (chan (chan broadcast)))
	sendc := make(chan interface{})
	go func() {
		currc := make(chan broadcast, 1)
		for {
			select {
			case v := <-sendc:
				c := make(chan broadcast, 1)
				b := broadcast{c: c, v: v}
				currc <- b
				currc = c
			case r := <-listenMouse:
				r <- currc
			case r := <-listenKeyboard:
				r <- currc
			}
		}
	}()
	return Broadcaster{
		ListenInt:    listenInt,
		ListenString: listenString,
		Sendc:        sendc,
	}
}

// start listening to the broadcasts.
func (b Broadcaster) Listen(sensor Sensor) {
	c := make(chan chan broadcast, 0)
	if tipo == 1 {
		b.ListenInt <- c
	} else if tipo == 2 {
		b.ListenString <- c
	}
	rv := Receiver{<-c, tipo}
	go rv.listen()
}

// broadcast a value to all listeners.
func (b Broadcaster) Write(v interface{}) { b.Sendc <- v }

type Receiver struct {
	C    chan broadcast
	Tipo int
}

func (r *Receiver) Read() interface{} {
	b := <-r.C
	v := b.v
	r.C <- b
	r.C = b.c
	return v
}

func (r *Receiver) listen() {
	for v := r.Read(); v != nil; v = r.Read() {
		if r.Tipo == 1 {
			fmt.Println(v)

		} else {
			fmt.Sprintf("Tipo excluido %d", r.Tipo)
		}
	}
}

func main() {
	var b = NewBroadcaster()
	b.Listen(1)
	b.Listen(2)
	b.Listen(2)

	for i := 0; i < 10; i++ {
		b.Write(fmt.Sprintf("Valor %d", i))
	}
	b.Write(nil)

	time.Sleep(3 * 1e9)
}
