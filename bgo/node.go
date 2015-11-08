package bgo

import "time"

type BaseNode struct {
	id          string
	name        string
	category    NodeCategorie
	title       string
	description string

}

func CreateBaseNode(title, desc string) *BaseNode {
	bn := &BaseNode{}
	bn.id = CreateUUID()
	bn.title = title
	bn.description = desc
	return bn
}


func (this *BaseNode) execute(tick Tick) Status {
	tick.enterNode(this);
	this.enter(tick)

	if (!tick.Blackboard.get("isOpen", tick.Tree.id, this.id)) {
		tick.openNode(this);
		tick.Blackboard.set("isOpen", true, tick.Tree.id, this.id)
		this.open(tick)

	}

	tick.tickNode(this)
	status := this.tick(tick)

	if status != RUNNING {
		tick.closeNode(this)
		tick.Blackboard.set("isOpen", false, tick.Tree.id, this.id)
		this.close(tick)
	}

	tick.exitNode(this)
	this.exit(tick)

	return status
}

func (bn *BaseNode) enter(tick Tick) {}
func (bn *BaseNode) exit(tick Tick) {}
func (bn *BaseNode) open(tick Tick) {}
func (bn *BaseNode) close(tick Tick) {}
func (bn *BaseNode) tick(tick Tick) Status { return ERROR}


/**
   * The Sequence node ticks its children sequentially until one of them
   * returns `FAILURE`, `RUNNING` or `ERROR`. If all children return the
   * success state, the sequence also returns `SUCCESS`.
  **/
type Sequence struct {
	*BaseNode
	children []*BaseNode
}

func (this *Sequence) tick(tick Tick) Status {
	for i := 0; i < len(this.children); i++ {
		status := this.children[i].execute(tick);

		if (status != SUCCESS) {
			return status;
		}
	}
	return SUCCESS;

}

func NewSequence(title string,children ...BaseNode) Sequence {
	sequence := &Sequence{}
	sequence.id = CreateUUID()
	sequence.category = COMPOSITE
	sequence.name = "Sequence"
	sequence.title = title
	sequence.description = "The Sequence node ticks its children sequentially until one of them returns `FAILURE`, `RUNNING` or `ERROR`"
	sequence.children = children
	return sequence
}


/**
   * Priority ticks its children sequentially until one of them returns
   * `SUCCESS`, `RUNNING` or `ERROR`. If all children return the failure state,
   * the priority also returns `FAILURE`.
**/
type Priority struct {
	*BaseNode
	children []*BaseNode
}

func (this *Priority) tick(tick Tick) Status {
	for i := 0; i < len(this.children); i++ {
		status := this.children[i].execute(tick);

		if (status != FAILURE) {
			return status;
		}
	}
	return FAILURE;

}

func NewPriority(title string,children ...BaseNode) Priority {
	priority := &Priority{}
	priority.id = CreateUUID()
	priority.category = COMPOSITE
	priority.name = "Priority"
	priority.title = title
	priority.description = "Priority ticks its children sequentially until one of them returns `SUCCESS`, `RUNNING` or `ERROR`"
	priority.children = children
	return priority
}

/**
   * Wait a few seconds.
  **/
type Wait struct {
	*BaseNode
	endTime chan time.Time
	delay	int16
}

func (this *Wait) open(tick Tick) {
	this.endTime = time.Tick(this.delay * time.Millisecond)
}

func (this *Wait) tick(tick Tick) Status {
	for range this.endTime {
		close(this.endTime)
		return SUCCESS
	}
	return RUNNING
}

func NewPWait(title string) Wait {
	wait := &Wait{}
	wait.id = CreateUUID()
	wait.category = ACTION
	wait.name = "Wait"
	wait.title = title
	wait.description = "Priority ticks its children sequentially until one of them returns `SUCCESS`, `RUNNING` or `ERROR`"
	return wait
}