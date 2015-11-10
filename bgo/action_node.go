package bgo

import "time"


/**
   * Wait a few seconds.
  **/
type Wait struct {
	BaseNode
	endTime <-chan time.Time
	delay	int64
}

func (this *Wait) Open(context *Context) {
	t := time.Duration(this.delay) * time.Millisecond
	this.endTime = time.Tick(t)
}

func (this *Wait) Tick(context Context) Status {
	for range this.endTime {
		return SUCCESS
	}
	return RUNNING
}

func NewWait(title string) *Wait {
	wait := &Wait{}
	wait.ID = CreateUUID()
	wait.Category = ACTION
	wait.Name = "Wait"
	wait.Title = title
	wait.Description = "Priority contexts its children sequentially until one of them returns `SUCCESS`, `RUNNING` or `ERROR`"
	return wait
}