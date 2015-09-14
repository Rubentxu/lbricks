package lbricks

import (
	"sync"

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

func (s System) Name() string {
	return s.name
}

func (s System) Priority() int {
	return 0
}


type RenderSystem struct {
	*System
	lock 	*sync.Mutex
	batch 	Batch
	views 	[]Renderable
	maxlen 	int
}

func NewRenderSystem(maxlen int, width, height float32, batch Batch,es EventSystem) *RenderSystem  {
	rs:= new(RenderSystem)
	rs.views = make([]Renderable,0, maxlen*2)
	rs.maxlen = maxlen
	rs.lock = new(sync.Mutex)
	rs.batch = batch
	rs.eventSystem = es
	return rs
}

func (rs *RenderSystem) AddView(view Renderable)  {
	rs.lock.Lock()
	if len(rs.views) == cap(rs.views) {
		// Reallocate
		rs.maxlen*=2
		re := make([]Renderable, 0, rs.maxlen)
		for _, v := range rs.views {
			re = append(re,v)
		}
		rs.views = re
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