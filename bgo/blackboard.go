package bgo

type TreeMemory struct {
	nodeMemory     map[string] interface{}
	openNodes      []BaseNode
	traversalDepth uint8
	traversalCycle uint8
}

type Blackboard struct {
	baseMemory map[string] interface{}
	treeMemory map[string] *TreeMemory
}

func CreateBlackboard() *Blackboard {
	return &Blackboard{
		baseMemory: make(map[string] interface{}),
		treeMemory: make(map[string]*TreeMemory),
	}
}

func (this *Blackboard) getTreeMemory(treeScope string)  *TreeMemory {
	elem, ok := this.treeMemory[treeScope]
	if ok {
		return elem
	} else {
		return &TreeMemory{
			nodeMemory: make(map[string] interface{}),
			openNodes: make([] BaseNode,0,30),
			traversalDepth: 0,
			traversalCycle: 0,
		}
	}
}

func  (this *Blackboard) getNodeMemory(treeMemory *TreeMemory, nodeScope string)  interface{} {
	memory := treeMemory.nodeMemory;
	elem, ok := memory[nodeScope]
	if ok {
		return elem
	} else {
		return treeMemory
	}
}

func  (this *Blackboard) getMemory(treeScope, nodeScope string)  interface{} {
	memory := this.baseMemory;
	if nodeScope !="" {
		memory2 := this.getTreeMemory(treeScope)
		if nodeScope !="" {
			return this.getNodeMemory(memory2,nodeScope)
		}
		return memory2
	}
	return memory
}


func (this  *Blackboard) get(key, treeScope, nodeScope string)  interface{} {
	memory := this.getMemory(treeScope,nodeScope)
	return memory[key]
}

func (this  *Blackboard) set(key string, value interface{}, treeScope, nodeScope string) {
	memory := this.getMemory(treeScope,nodeScope)
	memory[key] = value
}