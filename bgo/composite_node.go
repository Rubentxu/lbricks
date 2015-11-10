package bgo



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
		status := ExecuteNode(this.children[i],context)

		if (status != SUCCESS) {
			return status
		}
	}
	return SUCCESS

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
		status := ExecuteNode(this.children[i],context)

		if (status != FAILURE) {
			return status
		}
	}
	return FAILURE

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
