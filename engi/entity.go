ackage engi

import (
	"fmt"

	"github.com/trustmaster/goflow"
)

// COmentario
type Provider func(graph *flow.Graph) *Entity

type Entity struct {
	id    int
	tags  string
	logic flow.Graph
}

type EntityPool struct {
	entities  map[int]*Entity
	providers map[string]Provider
}

func CreatePool() *EntityPool {
	entities := make(map[int]*Entity)
	providers := make(map[string]Provider)
	return &EntityPool{entities, providers}
}

func (e *EntityPool) createEntity(name string) *Entity {
	graph := e.initGraphEntity()
	provider, ok := e.providers[name]
	if !ok {
		panic(fmt.Sprintf("%s entity does not exist", name))
	}
	return provider(graph)
}

func (e *EntityPool) initGraphEntity() *flow.Graph {
	n := new(flow.Graph) // creates the object in heap
	n.InitGraphState()   // allocates memory for the graph
	// Add processes to the network
	//n.Add(new(Greeter), "greeter")
	//n.Add(new(Printer), "printer")
	// Connect them with a channel
	n.Connect("greeter", "Res", "printer", "Line")
	// Our net has 1 inport mapped to greeter.Name
	n.MapInPort("In", "greeter", "Name")
	return n
}
