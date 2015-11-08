package main

import (
	"github.com/ajhager/engi"
	"math/rand"
	"github.com/Rubentxu/lbricks"
	"fmt"
)

var (
	bots   []*Bot
	on     bool
	num    int
	region *engi.Region
	batch  *engi.Batch
)

type Bot struct {
	*engi.Sprite
	DX, DY float32
}

type Game struct {
	*engi.Game
	EventSystem *lbricks.EventSystem
}

func (game *Game) Preload() {
	engi.Files.Add("bot", "data/icon.png")
}

func (game *Game) Setup() {
	engi.SetBg(0x2d3638)
	texture := engi.Files.Image("bot")
	region = engi.NewRegion(texture, 0, 0, int(texture.Width()), int(texture.Height()))
	batch = engi.NewBatch(800, 600)
}

var time float32

func (game *Game) Update(dt float32) {
	time += dt
	if time > 1 {
		println(int(engi.Time.Fps()))
		println(num)
		time = 0
	}

	if on {
		for i := 0; i < 10; i++ {
			bot := &Bot{engi.NewSprite(region, 0, 0), rand.Float32() * 500, rand.Float32()*500 - 250}
			bots = append(bots, bot)
		}
		num += 10
	}

	minX := float32(0)
	maxX := float32(800)
	minY := float32(0)
	maxY := float32(600)

	for _, bot := range bots {
		bot.Position.X += bot.DX * dt
		bot.Position.Y += bot.DY * dt
		bot.DY += 750 * dt

		if bot.Position.X < minX {
			bot.DX *= -1
			bot.Position.X = minX
		} else if bot.Position.X > maxX {
			bot.DX *= -1
			bot.Position.X = maxX
		}

		if bot.Position.Y < minY {
			bot.DY = 0
			bot.Position.Y = minY
		} else if bot.Position.Y > maxY {
			bot.DY *= -.85
			bot.Position.Y = maxY
			if rand.Float32() > 0.5 {
				bot.DY -= rand.Float32() * 200
			}
		}
	}
}

func (game *Game) Render() {
	batch.Begin()
	for _, bot := range bots {
		bot.Render(batch)
	}
	batch.End()
}

func (game *Game) Mouse(x, y float32, action engi.Action) {
	var mouseAction lbricks.MouseAction
	switch action {
	case engi.PRESS:
		mouseAction = lbricks.LEFT_BUTTON_DOWN
	case engi.RELEASE:
		mouseAction = lbricks.LEFT_BUTTON_UP
	}

	observers := game.EventSystem.EventPorts("MouseEvent")
	for _,el := range observers {
		el(lbricks.MouseEvent{x,y,mouseAction})
	}
}

func main() {
	game:= &Game{}
	game.EventSystem = lbricks.CreateEventSystem(100)

	signal := game.EventSystem.FromEvent("MouseEvent")

	signal.Filter(func(value interface{}) bool {
		return value.(lbricks.MouseEvent).Action == lbricks.LEFT_BUTTON_DOWN}).
	Subscribe(func (value interface {}) {
		on = true
		fmt.Println("From Subscribe True",value)
	})

	signal2 := game.EventSystem.FromEvent("MouseEvent")
	signal2.Filter(func(value interface{}) bool {
		return value.(lbricks.MouseEvent).Action == lbricks.LEFT_BUTTON_UP}).
	Subscribe(func (value interface {}) {
		on = false
		fmt.Println("From Subscribe False",value)

	})
	engi.Open("Botmark", 800, 600, false, game)
}
