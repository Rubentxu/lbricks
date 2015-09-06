package lbricks

import (
	"sync"
	"github.com/ajhager/engi"
)


type Systemer interface {
	Name() string
	Priority() int
	Dispose()
}



type System struct {
	name string
	eventSystem EventSystem
}

func (s System) Name()  {
	return s.name
}

func (s System) Priority() int {
	return 0
}


type RenderSystem struct {
	*System
	lock 	*sync.Mutex
	batch 	*engi.Batch
	views 	[]Renderer
	maxlen 	int
}

func NewRenderSystem(maxlen, width, height int, es EventSystem) *RenderSystem  {
	rs:= new(RenderSystem)
	rs.views = make([]Renderer,0, maxlen*2)
	rs.maxlen = maxlen
	rs.lock = new(sync.Mutex)
	rs.batch = engi.NewBatch(width, height)
	rs.eventSystem = es
	return rs
}

func (rs *RenderSystem) AddView(view Renderer)  {
	rs.lock.Lock()
	if len(rs.views) == cap(rs.views) {
		// Reallocate
		rs.maxlen*=1.1
		re := make([]View, 0, rs.maxlen)
		rs.views = append(re, rs.views)
	}
	rs.views = append(rs.views, view)
	rs.lock.Unlock()
}

func (rs *RenderSystem) Render() {
	rs.batch.Begin()
	for _, v := range rs.views {
		v.Render(rs.batch)
	}
	rs.batch.End()
	// eventSystem.PostRender()
}