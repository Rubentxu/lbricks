package lbricks

type Game struct {
	Pool 			EntityPool
	mouseChan 		[]chan *EventPacked
	keyboardChan 	[]chan *EventPacked
}

func (g *Game) RegisterInputChannel(eventType string, channel chan *EventPacked) {
	switch eventType {
	case "MouseEvent":
		append(mouseChan, channel)
	}
	case "KeyboardEvent":
		append(keyboardChan, channel)
	}

}

func (g *Game) Preload()                               {}
func (g *Game) Setup()                                 {}
func (g *Game) Close()                                 {}
func (g *Game) Update(dt float32)                      {}
func (g *Game) Step(step float64, numStep uint32)      {}
func (g *Game) Render()                                {}
func (g *Game) Resize(w, h int)                        {}

func (g *Game) Mouse(x, y float32, action MouseAction) {
	eventPacked := &EventPacked{1,&MouseEvent{x,y,0.0,action}}
	for _, e := range g.mouseChan {
    	e <- eventPacked
 	}
}

func (g *Game) Scroll(amount float32)                  {}

func (g *Game) Key(key Key, modifier Modifier, action KeyAction) {
	if key == Escape {
		Exit()
	}
	eventPacked := &EventPacked{1,&KeyboardEvent{key,modifier,action}}
	for _, e := range g.keyboardChan {
    	e <- eventPacked
 	}
}

func (g *Game) Type(char rune) {}
