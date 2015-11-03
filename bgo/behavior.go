package bgo

import (
	"math/rand"
)

type Status uint16
type NodeCategorie uint16

var (
	SUCCESS = Status(1)
	FAILURE = Status(2)
	RUNNING = Status(3)
	ERROR 	= Status(4)

	COMPOSITE = NodeCategorie(1)
	DECORATOR = NodeCategorie(2)
	ACTION	  = NodeCategorie(3)
	CONDITION = NodeCategorie(4)
)


func CreateUUID() string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	const (
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)
	b := make([]byte, 50)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := 49, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

type BehaviorTree struct {
	id   			string
	title			string
	description		string
	root 			BaseNode

}

func (bt *BehaviorTree) Tick(target interface{}, blackboard interface{})  {

}

func CreateBehaviorTree(title, desc string)  *BehaviorTree {
	bt := &BehaviorTree{}
	bt.id = CreateUUID()
	bt.title = title
	bt.description = desc
	return bt
}


type Blackboard struct {


}


type Tick struct {
	Tree  			BehaviorTree
	Target 			interface{}
	Blackboard		Blackboard
	openNodes		[] BaseNode
	nodeCount		int

}

func CreateTick()  *Tick {
	tick := &Tick{}
	tick.openNodes = make([]BaseNode,0,50)
	return tick
}

func (this Tick) enterNode(node BaseNode)  {
	this.nodeCount++
	this.openNodes = append(this.openNodes,node)
}

func (this Tick) closeNode(node BaseNode)  {
	this.openNodes =  this.openNodes[:len(this.openNodes)-1]
}



type BaseNode struct {
	id 				string
	name			string
	category		NodeCategorie
	title			string
	description 	string

}

func CreateBaseNode(title, desc string)  *BaseNode {
	bn := &BaseNode{}
	bn.id = CreateUUID()
	bn.title = title
	bn.description = desc
	return bn
}

func (bn *BaseNode) execute(tick Tick)  {

}