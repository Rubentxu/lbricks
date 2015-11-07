package lbricks
import "fmt"

type Event chan interface{}
type Predicate func(interface{}) bool
type Mapper func(interface{}) interface{}
type MultiMapper func(...interface{}) interface{}
type Reducer func(memo interface{}, element interface{}) interface{}
type Subscriber func(interface{})


type Signal struct {
	event Event
}

func (s Signal)  Map(fn Mapper) Signal {
	signal := Signal{make(Event)}

	go func() {
		for el := range s.event {
			signal.event <- fn(el)
		}
		fmt.Println("Close chan From Map")
		close(signal.event)
	}()
	return signal
}


func (s Signal) Filter(pred Predicate) Signal {
	signal := Signal{make(Event)}
	go func() {
		for el := range s.event {
			if keep := pred(el); keep {
				signal.event <- el
			}
		}
		fmt.Println("Close chan From filter")
		close(signal.event)
	}()
	return signal
}


func (s Signal) Reduce(red Reducer, memo interface{}) interface{} {
	for el := range s.event {
		memo = red(memo, el)
	}
	return memo
}


func (s Signal) Subscribe(fn Subscriber) {
	go func() {
		for el := range s.event {
			fn(el)
		}
	}()

}


func FromValues(els ... interface{}) Signal {
	c := make(Event)
	go func() {
		for _, el := range els {
			c <- el
		}
		fmt.Println("Close chan From value")
		close(c)
	}()
	return Signal{c}
}

