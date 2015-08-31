package lbricks

import (
	"fmt"
	"github.com/Rubentxu/lbricks/goflow"
)


type EntityProvider func(graphPool *GraphPool) *Entity
type GraphProvider func() (*flow.Graph)
type EntityID uint

// Entity
type Entity struct {
	id          EntityID
	logicGraphs map[string]flow.Graph
}

func NewEntity() *Entity {
	lg := make(map[string] flow.Graph)
	entity := &Entity{}
	entity.logicGraphs = lg
	return entity
}

func (e *Entity) Id() int {
	return e.id
}

func (e *Entity) HasLogicGraph(name string) bool {
	_, ok := e.logicGraphs[name]
	return ok
}

func (e *Entity) AddLogicGraph(name string, graph flow.Graph) bool {
	if !e.HasLogicGraph(name) {
		e.logicGraphs[name] = graph
		return true
	}
	return false
}

func (e *Entity) LogicGraph(name string) (*flow.Graph, bool) {
	elem, ok := e.logicGraphs[name]
	return elem, ok
}

func (e *Entity) GetPorts() [] *flow.Port {
	ports := []flow.Port{}
	for _, graph:= range e.logicGraphs {
		for _, p:= range graph.ListInPorts() {
			append(ports,p)
		}
	}
	return ports
}




// EntityPool
type EntityPool struct {
	idCount   uint64
	entities  map[int]*Entity
	unused    []Entity
	providers map[string]EntityProvider
	graphPool GraphPool
}

func CreateEntityPool(graphPool *GraphPool) *EntityPool {
	entities := make(map[int]*Entity)
	providers := make(map[string]EntityProvider)
	return &EntityPool{entities, providers, graphPool}
}

func (ep *EntityPool) addEntity(entity *Entity) {
	if _, ok := ep.entities[entity.Id()]; !ok {
		ep.entities[entity.Id()] = entity
	}
}

func (ep *EntityPool) CreateEntity(name string) []*flow.Port {
	provider, ok := ep.providers[name]
	if !ok {
		panic(fmt.Sprintf("%s entity does not exist", name))
		return nil
	}
	entity := provider(ep.graphPool)
	entity.id = ep.idCount + 1
	ep.addEntity(entity)
	return entity, entity.GetPorts()
}

func (ep *EntityPool) AddProvider(name string, provider EntityProvider) {
	if _, ok := ep.providers[name]; !ok {
		ep.providers[name] = provider
	}
}

func (ep *EntityPool) Entities()  {
	return ep.entities
}

func (ep *EntityPool) Entity(name string) (Entity, bool) {
	return  ep.entities[name]
}



// GraphPool
type GraphPool struct {
	providers map[string]GraphProvider
}

func CreateGraphPool() *GraphPool {
	providers := make(map[string]GraphProvider)
	return &GraphPool{providers}
}

func (gp *GraphPool) AddProvider(name string, provider GraphProvider) {
	if _, ok := gp.providers[name]; !ok {
		gp.providers[name] = provider
	}
}

func (gp *GraphPool) CreateLogicGraph(name string) *flow.Graph {
	provider, ok := gp.providers[name]
	if !ok {
		panic(fmt.Sprintf("%s graphLogic does not exist", name))
	}
	return provider()
}