package lbricks


type Game struct {
	Pool        EntityPool
	EventSystem EventSystem
}

func (g *Game) InitContext() {
	g.EventSystem = &EventSystem{}
	graphPool := CreateGraphPool()
	//graphPool.AddProvider("preloadGraph", preloadGraph)
	//graphPool.AddProvider("setupGraph", setupGraph)
	g.Pool = CreateEntityPool(graphPool)
}




