package lbricks


type StageProvider func(game GameContext) *Stage
type StageId uint


type Stage struct {
	id StageId
	name string
	game GameContext
}

type GameContext struct {
	Pool        EntityPool
	EventSystem EventSystem
	StageProviders map[string] StageProvider
	RenderSystem RenderSystem
}





