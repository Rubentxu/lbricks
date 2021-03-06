package bgo

type Context struct {
	*BehaviorTree
	ContextId 		string
	Target 			interface{}
	openNodes		map[string] Node
	nodeCount		int

}

func CreateContext(target interface{})  *Context {
	context := &Context{}
	context.ContextId = CreateUUID()
	context.Target = target
	context.openNodes = make(map[string]Node)
	return context
}

func (this Context) enterNode(node Node)  {
	this.nodeCount++
	this.openNodes[node.Id()] = node
}

func (this Context) exitNode(node Node) {/* TODO: call debug here*/}

func (this Context) openNode(node Node) {/* TODO: call debug here*/}

func (this Context) closeNode(node Node)  {
	delete(this.openNodes,node.Id())
}

func (this Context) contextNode(node Node) {/* TODO: call debug here*/}

func (this *Context) GetContextMemory() *Memory {
	return this.getExtendMemory(this.Id, this.ContextId)
}