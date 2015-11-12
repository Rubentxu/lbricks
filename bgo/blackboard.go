package bgo

import "time"

type Memory struct {
	Integer				map[string] 	int
	Float				map[string] 	float32
	Complex				map[string] 	complex64
	Byte				map[string] 	byte
	Rune 				map[string] 	rune
	Bool				map[string] 	bool
	String				map[string] 	string
	Memory				map[string] 	*Memory
	Time				map[string] 	time.Time
	ArrayInteger 		map[string] 	[]int
	ArrayFloat 			map[string] 	[]float32
	ArrayComplex 		map[string] 	[]complex64
	ArrayByte 			map[string] 	[]byte
	ArrayRune 			map[string] 	[]rune
	ArrayBool 			map[string] 	[]bool
	ArrayString 		map[string] 	[]string
	ArrayNode 			map[string] 	[]Node
}


func CreateMemory() *Memory {
	return &Memory{
		Integer			:		make(map[string] 	int),
		Float			:		make(map[string] 	float32),
		Complex			:		make(map[string] 	complex64),
		Byte			:		make(map[string] 	byte),
		Rune 			:		make(map[string] 	rune),
		Bool			:		make(map[string] 	bool),
		String			:		make(map[string] 	string),
		Memory			:		make(map[string] 	*Memory),
		Time			:		make(map[string] 	time.Time),
		ArrayInteger	:		make(map[string] 	[]int),
		ArrayFloat    	:		make(map[string] 	[]float32),
		ArrayComplex 	:		make(map[string] 	[]complex64),
		ArrayByte    	:		make(map[string] 	[]byte),
		ArrayRune    	:		make(map[string] 	[]rune),
		ArrayBool    	:		make(map[string] 	[]bool),
		ArrayString    	:		make(map[string] 	[]string),
		ArrayNode   :			make(map[string] 	[]Node),
	}
}

type Blackboard struct {
	memory *Memory
}

func CreateBlackboard() *Blackboard {
	return &Blackboard{
		memory: CreateMemory(),
	}
}

func (this *Blackboard) GetBaseMemory() *Memory {
	return this.memory
}

func (this *Blackboard) getMemory(memory *Memory,scope string) *Memory {
	ms, ok := memory.Memory[scope]
	if !ok {
		ms = CreateMemory()
		memory.Memory[scope] = ms
	}
	return ms
}

func (this *Blackboard) getTreeMemory(treeScope string) *Memory {
	return this.getMemory(this.memory,treeScope)
}

func  (this *Blackboard) getExtendMemory( treeScope, extendScope string) *Memory {
	treeMemory := this.getTreeMemory(treeScope)
	return this.getMemory(treeMemory,extendScope)
}
