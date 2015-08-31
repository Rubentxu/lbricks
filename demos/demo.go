package main

import (
	"fmt"
	"github.com/Rubentxu/lbricks/engi"
	"github.com/trustmaster/goflow"
	"github.com/Rubentxu/lbricks"
)

var (
	imprimir  string
	bot   engi.Drawable
	batch *engi.Batch
	font  *engi.Font
	tick  float64
)

func main() {
	game:= &Game{}
	game.InitContext()
	engi.Open("Demo", 800, 600, false, game.EventSystem)
}

type Game struct {
	lbricks.Game

}

func (g *Game) InitContext() {
	g.EventSystem = &lbricks.CreateEventSystem(100)
	graphPool := lbricks.CreateGraphPool()
	graphPool.AddProvider("DemoGraph",NewDemoGraphProvider)
	g.Pool = lbricks.CreateEntityPool(graphPool)
	g.Pool.AddProvider("DemoEntityProvider",DemoEntityProvider)

	
	g.EventSystem.RegisterInputChannel()
	ports := g.Pool.CreateEntity("DemoEntityProvider")
	for _, p := range ports {
		g.EventSystem.RegisterInputChannel(p.Port, p.Channel)
	}
}



func (g *Game) Preload() {
	engi.Files.Add("bot", "data/icon.png")
	engi.Files.Add("font", "data/font.png")
	batch = engi.NewBatch(engi.Width(), engi.Height())
}

func (g *Game) Setup() {
	engi.SetBg(0x2d3739)
	bot = engi.Files.Image("bot")
	font = engi.NewGridFont(engi.Files.Image("font"), 20, 20)
	tick = 1.0 / 40.0

	imprimir = "holasss"

	flow.RunNet(net)

}


func (g *Game) Render() {
	batch.Begin()
	font.Print(batch, imprimir, 10, 200, 0xffffff)
	batch.Draw(bot, 512, 320, 0.5, 0.5, 10, 10, 0, 0xffffff, 1)
	batch.End()
}

// A component that prints its input on screen
type Printer struct {
	flow.Component
	Line <-chan *engi.MouseAction // inport

}

// Prints a line when it gets it
func (p *Printer) OnLine(ms *lbricks.MouseEvent) {
	imprimir = fmt.Sprintf("Posicion del MouseX %.f MouseY %.f", ms.PosX, ms.PosY)
}

func DemoEntityProvider(pool lbricks.GraphPool) *lbricks.Entity{
	logicG := pool.CreateLogicGraph("DemoGraph")
	entity := lbricks.NewEntity()
	entity.AddLogicGraph("DemoGraph",logicG)
	return entity
}

func NewDemoGraphProvider() (*flow.Graph) {
	n := flow.NewGraph().(flow.Graph)

	msensor := lbricks.NewMouseSensor("mouseButtonUp", 1, engi.RIGHT_BUTTON_UP)
	n.Add(msensor, msensor.Name())
	n.Add(new(Printer), "printer")
	n.Connect(msensor.Name(), "Out", "printer", "Line")

	inMouseButtonUp := make([]chan *lbricks.MouseEvent)
	n.MapInPort("InMouseButtonUp", msensor.Name(), "In")
	n.SetInPort("InMouseButtonUp", inMouseButtonUp)
	return n
}
