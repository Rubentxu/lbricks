package engi

import (
	"fmt"

	"github.com/trustmaster/goflow"
)

type GraphProvider func(graph *flow.Graph) (*flow.Graph, map[string]chan *Event)

type GraphPool struct {
	providers map[string]GraphProvider
}

func CreateGraphPool() *GraphPool {
	providers := make(map[string]GraphProvider)
	return &GraphPool{providers}
}

func (e *GraphPool) CreateLogicGraph(name string) (*flow.Graph, map[string]chan *Event) {
	graph := new(flow.Graph)
	graph.InitGraphState()
	provider, ok := e.providers[name]
	if !ok {
		panic(fmt.Sprintf("%s graphLogic does not exist", name))
	}
	return provider(graph)
}

type EntityProvider func(graphPool GraphPool) *Entity

type EntityPool struct {
	entities  map[int]*Entity
	providers map[string]EntityProvider
	graphPool *GraphPool
}

type Entity struct {
	id          int
	tags        string
	logicGraphs map[string]*flow.Graph
}

func CreateEntityPool() *EntityPool {
	entities := make(map[int]*Entity)
	providers := make(map[string]EntityProvider)
	graphPool := CreateGraphPool()
	return &EntityPool{entities, providers, graphPool}
}

func (e *EntityPool) CreateEntity(name string) *Entity {
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
