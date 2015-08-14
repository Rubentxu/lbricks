package engi

import (
	"fmt"

	"github.com/trustmaster/goflow"
)

type Entity func(graph *flow.Graph) *flow.Graph

type EntityPool struct {
	entities        map[int]*flow.Graph
	entityFactories map[string]Entity
}

func CreatePool() *EntityPool {
	entities = make(map[int]*flow.Graph)
	entityFactories = make(map[string]Entity)
	return &EntityPool{entities, entityFactories}
}

func (e *EntityPool) createEntity(name string) *flow.Graph {
	graph := initEntity()
	entity, ok := entities[name]
	if !ok {
		panic(fmt.Sprintf("%s entity does not exist", name))
	}
	return entity(graph)
}

func (e *EntityPool) initEntity() *flow.Graph {
	n := new(flow.Graph) // creates the object in heap
	n.InitGraphState()   // allocates memory for the graph
	// Add processes to the network
	n.Add(new(Greeter), "greeter")
	n.Add(new(Printer), "printer")
	// Connect them with a channel
	n.Connect("greeter", "Res", "printer", "Line")
	// Our net has 1 inport mapped to greeter.Name
	n.MapInPort("In", "greeter", "Name")
	return n
}
