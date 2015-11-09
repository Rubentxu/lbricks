package bgo

import (
	"time"
	"fmt"
)

type Node interface {
	Id()		string
	Enter(context *Context)
	Exit(context *Context)
	Open(context *Context)
	Close(context *Context)
	Tick(context *Context) Status
}

type BaseNode struct {
	ID				 string
	Name 			string
	Category 		NodeCategorie
	Title 			string
	Description 	string

}

func CreateBaseNode(title string) *BaseNode {
	bn := &BaseNode{}
	bn.ID = CreateUUID()
	bn.Title = title
	return bn
}

func (this *BaseNode) Id() string{
	return this.ID
}


func (bn *BaseNode) Enter(context *Context) {}
func (bn *BaseNode) Exit(context *Context) {}
func (bn *BaseNode) Open(context *Context) {}
func (bn *BaseNode) Close(context *Context) {}
func (bn *BaseNode) Tick(context *Context) Status {
	fmt.Println("Test Context BaseNode  %s", bn.ID)
	return FAILURE
}

func Execute(this Node, context *Context) Status {
	context.enterNode(this);
	this.Enter(context)

	if _, ok := context.Blackboard.get("isOpen", context.Tree.Id, this.Id()); ok {
		context.openNode(this);
		context.Blackboard.set("isOpen", true, context.Tree.Id, this.Id())
		this.Open(context)

	}

	context.contextNode(this)
	status := this.Tick(context)

	if status != RUNNING {
		context.closeNode(this)
		context.Blackboard.set("isOpen", false, context.Tree.Id, this.Id())
		this.Close(context)
	}

	context.exitNode(this)
	this.Exit(context)

	return status
}


/**
   * The Sequence node contexts its children sequentially until one of them
   * returns `FAILURE`, `RUNNING` or `ERROR`. If all children return the
   * success state, the sequence also returns `SUCCESS`.
  **/
type Sequence struct {
	BaseNode
	children []Node
}

func (this *Sequence) Tick(context *Context) Status {
	for i := 0; i < len(this.children); i++ {
		status := Execute(this.children[i],context);

		if (status != SUCCESS) {
			return status;
		}
	}
	return SUCCESS;

}

func NewSequence(title string,children ...Node) *Sequence {
	sequence := &Sequence{}
	sequence.ID = CreateUUID()
	sequence.Category = COMPOSITE
	sequence.Name = "Sequence"
	sequence.Title = title
	sequence.Description = "The Sequence node contexts its children sequentially until one of them returns `FAILURE`, `RUNNING` or `ERROR`"
	sequence.children = children
	return sequence
}


/**
   * Priority contexts its children sequentially until one of them returns
   * `SUCCESS`, `RUNNING` or `ERROR`. If all children return the failure state,
   * the priority also returns `FAILURE`.
**/
type Priority struct {
	BaseNode
	children []Node
}

func (this *Priority) Tick(context *Context) Status {
	for i := 0; i < len(this.children); i++ {
		status := Execute(this.children[i],context);

		if (status != FAILURE) {
			return status;
		}
	}
	return FAILURE;

}

func NewPriority(title string,children ...Node) *Priority {
	priority := &Priority{}
	priority.ID = CreateUUID()
	priority.Category = COMPOSITE
	priority.Name = "Priority"
	priority.Title = title
	priority.Description = "Priority contexts its children sequentially until one of them returns `SUCCESS`, `RUNNING` or `ERROR`"
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