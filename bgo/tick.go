package bgo

type Tick struct {
	Tree  			BehaviorTree
	Target 			interface{}
	Blackboard		Blackboard
	openNodes		[] BaseNode
	nodeCount		int

}

func CreateTick(target interface{}, blackboard Blackboard)  *Tick {
	tick := &Tick{}
	tick.Target = target
	tick.Blackboard = blackboard
	tick.openNodes = make([]BaseNode,0,50)
	return tick
}

func (this Tick) enterNode(node BaseNode)  {
	this.nodeCount++
	this.openNodes = append(this.openNodes,node)
}

func (this Tick) exitNode(node BaseNode) {/* TODO: call debug here*/}

func (this Tick) openNode(node BaseNode) {/* TODO: call debug here*/}

func (this Tick) closeNode(node BaseNode)  {
	this.openNodes =  this.openNodes[:len(this.openNodes)-1]
}

func (this Tick) tickNode(node BaseNode) {/* TODO: call debug here*/}