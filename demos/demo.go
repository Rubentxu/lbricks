package main

import (
	"fmt"

	"github.com/Rubentxu/lbricks"
	"github.com/Rubentxu/lbricks/engi"
	"github.com/trustmaster/goflow"
)

var (
	imprimir  string
)

type Game struct {
	*lbricks.Game
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
	game.Pool := CreateEntityPool()
	game.Pool.AddProvider("DemoEntity", DemoEntityProvider)
	game.Pool.AddProvider(name string, provider GraphProvider)
	imprimir = "holasss"

	flow.RunNet(net)

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
	game:= &Game{}
	game.InitContext()
	engi.Open("Demo", 800, 600, false, game.EventSystem)
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

func DemoEntityProvider(pool GraphPool) (*Entity, map[string]chan *EventPacked){

}

func NewDemoGraphProvider() (*flow.Graph, map[string]chan *EventPacked) {
	n := new(flow.Graph) // creates the object in heap
	n.InitGraphState()    // allocates me(*flow.Graph, map[string]chan *EventPacked)mory for the graph

	msensor := NewMouseSensor("mouseButtonUp", 1, engi.RIGHT_BUTTON_UP)
	n.Add(msensor, "mouseSensor")
	n.Add(new(Printer), "printer")
	n.Connect("mouseSensor", "Out", "printer", "Line")

	inputs := make(map(string)chan *EventPacked)
	inputs["InMouseButtonUp"]= make(chan *RequestPacket)
	n.MapInPort("InMouseButtonUp", "mouseSensor", "In")
	net.SetInPort("InMouseButtonUp", inputs["InMouseButtonUp"])
	return n,inputs

}
