package lbricks


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

}