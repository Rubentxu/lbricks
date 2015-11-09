package bgo

type Tick struct {
	Tree  			*BehaviorTree
	Target 			interface{}
	Blackboard		*Blackboard
	openNodes		map[string] Node
	nodeCount		int

}

func CreateTick(target interface{}, blackboard *Blackboard)  *Tick {
	tick := &Tick{}
	tick.Target = target
	tick.Blackboard = blackboard
	tick.openNodes = make(map[string]Node)
	return tick
}

func (this Tick) enterNode(node Node)  {
	this.nodeCount++
	this.openNodes[node.Id()] = node
}

func (this Tick) exitNode(node Node) {/* TODO: call debug here*/}

func (this Tick) openNode(node Node) {/* TODO: call debug here*/}

func (this Tick) closeNode(node Node)  {
	delete(this.openNodes,node.Id())
}

func (this Tick) tickNode(node Node) {/* TODO: call debug here*/}