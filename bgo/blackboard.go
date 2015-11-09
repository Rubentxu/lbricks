package bgo

type Blackboard struct {
	baseMemory map[string] interface{}
	treeMemory map[string] map[string] interface{}
}

func CreateBlackboard() *Blackboard {
	return &Blackboard{
		baseMemory: make(map[string] interface{}),
		treeMemory: make(map[string] map[string] interface{}),
	}
}

func (this *Blackboard) getTreeMemory(treeScope string)  map[string] interface{} {
	elem, ok := this.treeMemory[treeScope]
	if !ok {
		treeMemory := make(map[string] interface{})
		treeMemory["nodeMemory"] = make(map[string] map[string] interface{})
		treeMemory["openNodes"] = make(map[string] Node)
		treeMemory["traversalDepth"] = 0
		treeMemory["traversalCycle"] = 0
		this.treeMemory[treeScope] = treeMemory
		elem = treeMemory
	}
	return elem
}

func  (this *Blackboard) getNodeMemory(treeMemory map[string] interface{}, nodeScope string)  ( map[string] interface{},bool) {
	memory := treeMemory["nodeMemory"].(map[string] map[string] interface{})
	elem, ok := memory[nodeScope]
	return elem, ok
}

func  (this *Blackboard) getMemory(treeScope, nodeScope string)  map[string] interface{} {
	memory := this.baseMemory;
	if treeScope !="" {
		memory := this.getTreeMemory(treeScope)
		if nodeScope !="" {
			elem, ok :=this.getNodeMemory(memory,nodeScope)
			if ok {
				memory = elem
			}
		}
	}
	return memory
}


func (this  *Blackboard) get(key, treeScope, nodeScope string)  (interface{}, bool) {
	memory := this.getMemory(treeScope,nodeScope)
	elem, ok := memory[key]
	return elem, ok

}

func (this  *Blackboard) set(key string, value interface{}, treeScope, nodeScope string)  {
	memory := this.getMemory(treeScope,nodeScope)
	memory[key] = value
}