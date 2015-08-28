package lbricks


type EventSystem struct {
	mouseChan 		[]chan *EventPacked
	keyboardChan 	[]chan *EventPacked
}

func (g *EventSystem) RegisterInputChannel(eventType string, channel chan *EventPacked) {
	switch eventType {
	case "MouseEvent":
		append(mouseChan, channel)
	}
	case "KeyboardEvent":
		append(keyboardChan, channel)
	}

}

func (g *EventSystem) Preload()                               {}
func (g *EventSystem) Setup()                                 {}
func (g *EventSystem) Close()                                 {}
func (g *EventSystem) Update(dt float32)                      {}
func (g *EventSystem) Step(step float64, numStep uint32)      {}
func (g *EventSystem) Render()                                {}
func (g *EventSystem) Resize(w, h int)                        {}

func (g *EventSystem) Mouse(x, y float32, action MouseAction) {
	eventPacked := &EventPacked{1,&MouseEvent{x,y,0.0,action}}
	for _, e := range g.mouseChan {
    	e <- eventPacked
 	}
}

func (g *EventSystem) Scroll(amount float32)                  {}

func (g *EventSystem) Key(key Key, modifier Modifier, action KeyAction) {
	if key == Escape {
		Exit()
	}
	eventPacked := &EventPacked{1,&KeyboardEvent{key,modifier,action}}
	for _, e := range g.keyboardChan {
    	e <- eventPacked
 	}
}

func (g *EventSystem) Type(char rune) {}
