package lbricks

type Game struct {
	Pool        EntityPool
	EventSystem Responder
}

func (g *Game) InitContext() {
	g.EventSystem := &EventSystem{}
	graphPool := CreateGraphPool()
	g.EntityPool = CreateEntityPool(graphPool)
}
