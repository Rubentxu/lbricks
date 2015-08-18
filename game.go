package lbricks

type Game struct {
	pool *EntityPool
}

func (g *Game) Preload()                               {}
func (g *Game) Setup()                                 {}
func (g *Game) Close()                                 {}
func (g *Game) Update(dt float32)                      {}
func (g *Game) Step(step float64, numStep uint32)      {}
func (g *Game) Render()                                {}
func (g *Game) Resize(w, h int)                        {}
func (g *Game) Mouse(x, y float32, action MouseAction) {}
func (g *Game) Scroll(amount float32)                  {}
func (g *Game) Key(key Key, modifier Modifier, action KeyAction) {
	if key == Escape {
		Exit()
	}
}
func (g *Game) Type(char rune) {}
