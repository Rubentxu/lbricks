package main

import (
	"fmt"

	"github.com/Rubentxu/lbricks"
	"github.com/Rubentxu/lbricks/engi"
	"github.com/trustmaster/goflow"
)

var (
	imprimir  string
	mouseChan chan *lbricks.MouseSignal
)

type Game struct {
	*engi.Game
	bot   engi.Drawable
	batch *engi.Batch
	font  *engi.Font
	tick  float64
}

func (game *Game) Preload() {
	engi.Files.Add("bot", "data/icon.png")
	engi.Files.Add("font", "data/font.png")
	game.batch = engi.NewBatch(engi.Width(), engi.Height())
}

func (game *Game) Setup() {
	engi.SetBg(0x2d3739)
	game.bot = engi.Files.Image("bot")
	game.font = engi.NewGridFont(engi.Files.Image("font"), 20, 20)
	game.tick = 1.0 / 40.0
	net := NewGreetingApp()
	imprimir = "holasss"
	// We need a channel to talk to it
	mouseChan = make(chan *lbricks.MouseSignal)
	net.SetInPort("In", mouseChan)
	// Run the net
	flow.RunNet(net)
	// Now we can send some names and see what happens

}

func (g *Game) Update(dt float32) {

}

func (game *Game) Render() {
	game.batch.Begin()
	game.font.Print(game.batch, imprimir, 10, 200, 0xffffff)
	game.batch.Draw(game.bot, 512, 320, 0.5, 0.5, 10, 10, 0, 0xffffff, 1)
	game.batch.End()
}

func main() {
	engi.Open("Hello", 1024, 640, false, &Game{})
}

func (game *Game) Mouse(x, y float32, event engi.MouseEvent) {
	var ms = &lbricks.MouseSignal{x, y, 0, event}
	mouseChan <- ms
}

// A component that prints its input on screen
type Printer struct {
	flow.Component
	Line <-chan *lbricks.MouseSignal // inport

}

// Prints a line when it gets it
func (p *Printer) OnLine(ms *lbricks.MouseSignal) {
	imprimir = fmt.Sprintf("Posicion del MouseX %.f MouseY %.f", ms.PosX, ms.PosY)
}

// Our greeting network
type GreetingApp struct {
	flow.Graph // graph "superclass" embedded
}

// Graph constructor and structure definition
func NewGreetingApp() *GreetingApp {
	n := new(GreetingApp) // creates the object in heap
	n.InitGraphState()    // allocates memory for the graph
	// Add processes to the network
	msensor := lbricks.MouseSensor{Event: engi.RIGHT_BUTTON_UP}
	n.Add(&msensor, "mouseSensor")
	n.Add(new(Printer), "printer")
	// Connect them with a channel
	n.Connect("mouseSensor", "Out", "printer", "Line")
	// Our net has 1 inport mapped to greeter.Name
	n.MapInPort("In", "mouseSensor", "In")
	return n
}
