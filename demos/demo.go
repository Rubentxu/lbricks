package main
import (
	"github.com/Rubentxu/lbricks/engi"
	"github.com/Rubentxu/lbricks"
)



var (
	imprimir  string
	bot   lbricks.Drawable
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
	EventSystem lbricks.EventSystem

}

func (g *Game) InitContext() {
	g.EventSystem = &lbricks.CreateEventSystem(100)

}



func  Preload(input chan *lbricks.PreloadEvent) {
	for {
		<- input
		engi.Files.Add("bot", "data/icon.png")
		engi.Files.Add("font", "data/font.png")
		batch = engi.NewBatch(engi.Width(), engi.Height())
	}
}

func Setup(input chan *lbricks.SetupEvent) {
	<- input
	engi.SetBg(0x2d3739)
	bot = engi.Files.Image("bot")
	font = engi.NewGridFont(engi.Files.Image("font"), 20, 20)
	tick = 1.0 / 40.0

	imprimir = "holasss"

}


func (g *Game) Render() {
	batch.Begin()
	font.Print(batch, imprimir, 10, 200, 0xffffff)
	batch.Draw(bot, 512, 320, 0.5, 0.5, 10, 10, 0, 0xffffff, 1)
	batch.End()
}

