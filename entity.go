package lbricks

import (
	"fmt"
	"github.com/trustmaster/goflow"
)


type EntityProvider func(graphPool *GraphPool) (*Entity, map[string]chan *EventPacket)
type GraphProvider func() (*flow.Graph, map[string]chan *EventPacket)
//type ComponentProvider func() interface{}
//type SensorProvider func(graph *flow.Graph) *Sensor


type EntityID uint

// Entity
type Entity struct {
	id          EntityID
	logicGraphs map[string]flow.Graph
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

func (e *Entity) GetLogicGraph(name string) (*flow.Graph, bool) {
	elem, ok := e.logicGraphs[name]
	return elem, ok
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

func (ep *EntityPool) CreateEntity(name string) (*Entity, map[string]chan *EventPacket) {
	provider, ok := ep.providers[name]
	if !ok {
		panic(fmt.Sprintf("%s entity does not exist", name))
		return nil
	}
	entity, inputs := provider(ep.graphPool)
	entity.id = ep.idCount + 1
	ep.AddEntity(entity)
	return entity, inputs
}

func (ep *EntityPool) AddEntity(entity *Entity) {
	if _, ok := ep.entities[entity.Id()]; !ok {
		ep.entities[entity.Id()] = entity
	}
}

func (ep *EntityPool) AddProvider(name string, provider EntityProvider) {
	if _, ok := ep.providers[name]; !ok {
		ep.providers[name] = provider
	}
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

func (gp *GraphPool) CreateLogicGraph(name string) (*flow.Graph, map[string]chan *EventPacket) {
	provider, ok := gp.providers[name]
	if !ok {
		panic(fmt.Sprintf("%s graphLogic does not exist", name))
	}
	return provider()
}