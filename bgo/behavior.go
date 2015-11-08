package bgo

import (
	"math/rand"
	"math"
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

func (this *BehaviorTree) Tick(target interface{}, blackboard Blackboard)  {
	tick := CreateTick(target,blackboard)
	tick.Tree = this

	state := this.root.execute(tick)

	lastOpenNodes := blackboard.get("openNodes", this.id, nil).([]BaseNode)
	currOpenNodes := tick.openNodes.([]BaseNode)

	start := 0

	for i :=0; i < math.Min(len(lastOpenNodes), len(currOpenNodes)); i++ {
		start = i+1
		if lastOpenNodes[i] != currOpenNodes[i] {
			break
		}
	}

	for i := len(lastOpenNodes) - 1; i >= start; i-- {
		lastOpenNodes[i].close(tick)
	}

	blackboard.set("openNodes",currOpenNodes, this.id, nil)
	blackboard.set("nodeCount",tick.nodeCount, this.id, nil)

	return state

}

func CreateBehaviorTree(title, desc string)  *BehaviorTree {
	bt := &BehaviorTree{}
	bt.id = CreateUUID()
	bt.title = title
	bt.description = desc
	return bt
}


