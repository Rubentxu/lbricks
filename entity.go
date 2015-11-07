package lbricks
import "fmt"


type EntityProvider func() *Entity

type EntityID uint

// Entity
type Entity struct {
	id          EntityID

}

func NewEntity() *Entity {
	entity := &Entity{}
	return entity
}

func (e *Entity) Id() EntityID {
	return e.id
}


// EntityPool
type EntityPool struct {
	idCount   uint64
	entities  map[EntityID]*Entity
	unused    []Entity
	providers map[string]EntityProvider

}

func CreateEntityPool() *EntityPool {
	e := make(map[EntityID]*Entity)
	p := make(map[string]EntityProvider)
	return &EntityPool{entities: e, providers: p}
}

func (ep *EntityPool) addEntity(entity *Entity) {
	if _, ok := ep.entities[entity.Id()]; !ok {
		ep.entities[entity.Id()] = entity
	}
}

func (ep *EntityPool) CreateEntity(name string)  *Entity{
	provider, ok := ep.providers[name]
	if !ok {
		panic(fmt.Sprintf("%s entity does not exist", name))
		return nil
	}
	entity := provider()
	entity.id = EntityID(int(ep.idCount)+1)
	ep.addEntity(entity)
	return entity
}

func (ep *EntityPool) AddProvider(name string, provider EntityProvider) {
	if _, ok := ep.providers[name]; !ok {
		ep.providers[name] = provider
	}
}

func (ep *EntityPool) Entities() map[EntityID]*Entity{
	return ep.entities
}

func (ep *EntityPool) Entity(id EntityID) (*Entity, bool) {
	elem,ok :=  ep.entities[id]
	return elem,ok
}