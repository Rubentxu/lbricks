package main

import (
	"fmt"
	"github.com/ajhager/engi"
	"github.com/trustmaster/goflow"
)

var (
	mouseChan chan *MouseSignal
	imprimir  string
)

type MouseSignal struct {
	posX, posY float32
	action     engi.Action
}

type Game struct {
	*engi.Game
	bot   engi.Drawable
	batch *engi.Batch
	font  *engi.Font
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
	net := NewGreetingApp()
	imprimir = "holasss"
	// We need a channel to talk to it
	mouseChan = make(chan *MouseSignal)
	net.SetInPort("In", mouseChan)
	// Run the net
	flow.RunNet(net)
	// Now we can send some names and see what happens

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

func (game *Game) Mouse(x, y float32, action engi.Action) {
	var ms = &MouseSignal{x, y, action}
	mouseChan <- ms
}

// A component that generates greetings
type MouseSensor struct {
	flow.Component                     // component "superclass" embedded
	Signal         <-chan *MouseSignal // input port
	ResSignal      chan<- *MouseSignal // output port
}

func (g *MouseSensor) OnSignal(ms *MouseSignal) {
	if ms.action == engi.PRESS {
		g.ResSignal <- ms
	}
}

// A component that prints its input on screen
type Printer struct {
	flow.Component
	Line <-chan *MouseSignal // inport

}

// Prints a line when it gets it
func (p *Printer) OnLine(ms *MouseSignal) {
	imprimir = fmt.Sprintf("Posicion del MouseX %.f MouseY %.f", ms.posX, ms.posY)
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
	n.Add(new(MouseSensor), "mouseSensor")
	n.Add(new(Printer), "printer")
	// Connect them with a channel
	n.Connect("mouseSensor", "ResSignal", "printer", "Line")
	// Our net has 1 inport mapped to greeter.Name
	n.MapInPort("In", "mouseSensor", "Signal")
	return n
}
