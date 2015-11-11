package bgo

type Node interface {
	Id()		string
	Enter(context *Context)
	Exit(context *Context)
	Open(context *Context)
	Close(context *Context)
	Tick(context *Context) Status
}

type BaseNode struct {
	ID				string
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
	return FAILURE
}

func ExecuteNode(node Node, context *Context) Status {
	nodeMemory := context.GetNodeMemory(node)
	context.enterNode(node);
	node.Enter(context)

	if _, ok := nodeMemory.Bool["isOpen"]; !ok {
		context.openNode(node);
		nodeMemory.Bool["isOpen"] =  true
		node.Open(context)

	}

	context.contextNode(node)
	status := node.Tick(context)

	if status != RUNNING {
		context.closeNode(node)
		nodeMemory.Bool["isOpen"]= false
		node.Close(context)
	}

	context.exitNode(node)
	node.Exit(context)

	return status
}

