package bgo
import "math/rand"



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
	Id 				string
	Title			string
	Description 	string
	Root 			Node

}

func (this *BehaviorTree) Tick(target interface{}, blackboard *Blackboard) Status {
	tick := CreateContext(target,blackboard)
	tick.Tree = this

	state := Execute(this.Root,tick)

	var lastOpenNodes map[string] Node
	if nodes, ok := blackboard.get("openNodes", this.Id, "");ok {
		lastOpenNodes = nodes.(map[string] Node)
	}
	currOpenNodes := tick.openNodes


	for _,lastNode := range lastOpenNodes {
		for _,currNode := range currOpenNodes {
			if lastNode.Id() != currNode.Id() {
				lastNode.Close(tick)
			}
		}
	}

	blackboard.set("openNodes",currOpenNodes, this.Id, "")
	blackboard.set("nodeCount",tick.nodeCount, this.Id, "")

	return state

}



func CreateBehaviorTree(title, desc string)  *BehaviorTree {
	bt := &BehaviorTree{}
	bt.Id = CreateUUID()
	bt.Title = title
	bt.Description = desc
	return bt
}


