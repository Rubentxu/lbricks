package lbricks


type EventSystem struct {
	eventPorts map[string] [] func(interface{})
}


func CreateEventSystem(capacity int) *EventSystem  {
	eventSystem := new(EventSystem)
	eventSystem.eventPorts = make(map[string] [] func(interface{}))
	return eventSystem
}


func  (es *EventSystem) FromEvent(eventType string) Signal {
	c := make(chan interface{})
	es.addEventListener(eventType, func (el interface{}) { c <- el	})
	return Signal{c}
}


func (es *EventSystem) addEventListener(eventType string, fn func(el interface{}))  {
	es.eventPorts[eventType] = append(es.EventPorts(eventType),fn)
}


func (es *EventSystem) EventPorts(eventType string) [] func(interface{}) {
	if _,ok := es.eventPorts[eventType]; !ok {
		es.eventPorts[eventType] =  make([]func(interface{}),0,1)
	}
	return es.eventPorts[eventType]
}

