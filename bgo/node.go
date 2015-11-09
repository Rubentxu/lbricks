package bgo

import "time"

type Node interface {
	Id() 				string
	execute(tick *Tick) Status
	enter(tick *Tick)
	open(tick *Tick)
	close(tick *Tick)
	tick(tick *Tick) Status
}

type BaseNode struct {
	id 				string
	Name 			string
	Category 		NodeCategorie
	Title 			string
	Description 	string

}

func CreateBaseNode(title, desc string) *BaseNode {
	bn := &BaseNode{}
	bn.id = CreateUUID()
	bn.Title = title
	bn.Description = desc
	return bn
}

func (this *BaseNode) Id() string{
	return this.id
}


func (this *BaseNode) SetId(id string) {
	 this.id = id
}

func (this *BaseNode) execute(tick *Tick) Status {
	tick.enterNode(this);
	this.enter(tick)

	if _, ok := tick.Blackboard.get("isOpen", tick.Tree.Id, this.id); ok {
		tick.openNode(this);
		tick.Blackboard.set("isOpen", true, tick.Tree.Id, this.id)
		this.open(tick)

	}

	tick.tickNode(this)
	status := this.tick(tick)

	if status != RUNNING {
		tick.closeNode(this)
		tick.Blackboard.set("isOpen", false, tick.Tree.Id, this.id)
		this.close(tick)
	}

	tick.exitNode(this)
	this.exit(tick)

	return status
}

func (bn *BaseNode) enter(tick *Tick) {}
func (bn *BaseNode) exit(tick *Tick) {}
func (bn *BaseNode) open(tick *Tick) {}
func (bn *BaseNode) close(tick *Tick) {}
func (bn *BaseNode) tick(tick *Tick) Status { return ERROR}


/**
   * The Sequence node ticks its children sequentially until one of them
   * returns `FAILURE`, `RUNNING` or `ERROR`. If all children return the
   * success state, the sequence also returns `SUCCESS`.
  **/
type Sequence struct {
	BaseNode
	children []Node
}

func (this *Sequence) tick(tick *Tick) Status {
	for i := 0; i < len(this.children); i++ {
		status := this.children[i].execute(tick);

		if (status != SUCCESS) {
			return status;
		}
	}
	return SUCCESS;

}

func NewSequence(title string,children ...Node) *Sequence {
	sequence := &Sequence{}
	sequence.id = CreateUUID()
	sequence.Category = COMPOSITE
	sequence.Name = "Sequence"
	sequence.Title = title
	sequence.Description = "The Sequence node ticks its children sequentially until one of them returns `FAILURE`, `RUNNING` or `ERROR`"
	sequence.children = children
	return sequence
}


/**
   * Priority ticks its children sequentially until one of them returns
   * `SUCCESS`, `RUNNING` or `ERROR`. If all children return the failure state,
   * the priority also returns `FAILURE`.
**/
type Priority struct {
	BaseNode
	children []Node
}

func (this *Priority) tick(tick *Tick) Status {
	for i := 0; i < len(this.children); i++ {
		status := this.children[i].execute(tick);

		if (status != FAILURE) {
			return status;
		}
	}
	return FAILURE;

}

func NewPriority(title string,children ...Node) *Priority {
	priority := &Priority{}
	priority.id = CreateUUID()
	priority.Category = COMPOSITE
	priority.Name = "Priority"
	priority.Title = title
	priority.Description = "Priority ticks its children sequentially until one of them returns `SUCCESS`, `RUNNING` or `ERROR`"
	priority.children = children
	return priority
}

/**
   * Wait a few seconds.
  **/
type Wait struct {
	BaseNode
	endTime <-chan time.Time
	delay	int64
}

func (this *Wait) open(tick *Tick) {
	t := time.Duration(this.delay) * time.Millisecond
	this.endTime = time.Tick(t)
}

func (this *Wait) tick(tick Tick) Status {
	for range this.endTime {
		return SUCCESS
	}
	return RUNNING
}

func NewPWait(title string) *Wait {
	wait := &Wait{}
	wait.id = CreateUUID()
	wait.Category = ACTION
	wait.Name = "Wait"
	wait.Title = title
	wait.Description = "Priority ticks its children sequentially until one of them returns `SUCCESS`, `RUNNING` or `ERROR`"
	return wait
}