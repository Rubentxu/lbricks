package lbricks

type Game struct {
	Pool        EntityPool
	EventSystem Responder
}

func (g *Game) InitContext() {
	g.EventSystem := &EventSystem{}
	graphPool := CreateGraphPool()
	graphPool.AddProvider("preloadGraph", preloadGraph)
	graphPool.AddProvider("setupGraph", setupGraph)
	g.EntityPool = CreateEntityPool(graphPool)
}


func preloadGraph() (*flow.Graph, map[string]chan *EventPacked){
	
}



func setupGraph() (*flow.Graph, map[string]chan *EventPacked){
	
}


