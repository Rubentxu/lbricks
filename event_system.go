package lbricks


type EventSystem struct {
	eventPort map[string] [] func(interface{})
}


func CreateEventSystem(capacity int) *EventSystem  {
	eventSystem := new(EventSystem)
	return eventSystem
}


func  (es *EventSystem) FromEvent(eventType string) Signal {
	c := make(Event)
	defer close(c)
	es.addEventListener(eventType, func (el interface{}) { c <- el	})
	return Signal{c}
}


func (es *EventSystem) addEventListener(eventType string, fn func(el interface{}))  {
	if _,ok := es.eventPort[eventType]; !ok {
		es.eventPort[eventType] =  make([]func(interface{}))
	}
	es.eventPort[eventType] = append(es.eventPort[eventType],fn)
}


func (es *EventSystem) EventPort() map[string] [] func(interface{}) {
	return es.eventPort
}

