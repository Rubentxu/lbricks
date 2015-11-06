package funtional

type Event chan interface{}
type Predicate func (interface{}) bool
type Mapper func (interface{}) interface{}
type MultiMapper func (...interface{}) interface{}
type Reducer func (memo interface{}, element interface{}) interface{}


type Signal struct {
	event Event
}

func (s *Signal)  Map(fn Mapper) Signal {
	c := make(Event)

	go func () {
		for el := range s.event {
			c <- fn(el)
		}
		close(c)
	}()
	return Signal{c}
}


func (s *Signal) Filter(pred Predicate) Signal {
	c := make(Event)
	go func () {
		for el := range s.event {
			if keep := pred(el); keep {
				c <- el
			}
		}
		close(c)
	}()
	return Signal{c}
}


func  prueba(s Signal) Signal {

	return s.Map(func (interface{})).Filter().Map()
}