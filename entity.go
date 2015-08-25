package lbricks

import (
	"fmt"
	"github.com/trustmaster/goflow"
)


type EntityProvider func(graphPool GraphPool) (*Entity, map[string]chan *EventPacked)
type GraphProvider func() (*flow.Graph, map[string]chan *EventPacked)
//type ComponentProvider func() interface{}
//type SensorProvider func(graph *flow.Graph) *Sensor




// Entity
type Entity struct {
	id          int
	Tags        string
	logicGraphs map[string]*flow.Graph
}

func (e *Entity) Id() int {
	return e.id
}




// EntityPool
type EntityPool struct {
	idCount   uint64
	entities  map[int]*Entity
	providers map[string]EntityProvider
	graphPool *GraphPool
}

func CreateEntityPool() *EntityPool {
	entities := make(map[int]*Entity)
	providers := make(map[string]EntityProvider)
	graphPool := CreateGraphPool()
	return &EntityPool{entities, providers, graphPool}
}

func (ep *EntityPool) CreateEntity(name string) *Entity {
	provider, ok := ep.providers[name]
	if !ok {
		panic(fmt.Sprintf("%s entity does not exist", name))
	}
	entity, sensors := provider(ep.graphPool)
	entity.id = ep.idCount++
	ep.AddEntity(entity)
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

func (gp *GraphPool) CreateLogicGraph(name string) (*flow.Graph, map[string]chan *EventPacked) {
	provider, ok := gp.providers[name]
	if !ok {
		panic(fmt.Sprintf("%s graphLogic does not exist", name))
	}
	return provider()
}