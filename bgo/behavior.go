package bgo
import "math/rand"



type Status string
type NodeCategorie string

var (
	SUCCESS = Status("Succes")
	FAILURE = Status("Failure")
	RUNNING = Status("Runnig")
	ERROR 	= Status("Error")

	COMPOSITE = NodeCategorie("Composite")
	DECORATOR = NodeCategorie("Decorator")
	ACTION	  = NodeCategorie("Action")
	CONDITION = NodeCategorie("Condition")
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

func (this *BehaviorTree) Tick(context *Context) Status {
	context.Tree = this

	state := ExecuteNode(this.Root,context)

	var lastOpenNodes map[string] Node
	if nodes, ok := context.Blackboard.Get("openNodes", this.Id, "");ok {
		lastOpenNodes = nodes.(map[string] Node)
	}
	currOpenNodes := context.openNodes


	for _,lastNode := range lastOpenNodes {
		for _,currNode := range currOpenNodes {
			if lastNode.Id() != currNode.Id() {
				lastNode.Close(context)
			}
		}
	}

	context.Blackboard.Set("openNodes",currOpenNodes, this.Id, "")
	context.Blackboard.Set("nodeCount",context.nodeCount, this.Id, "")

	return state

}



func CreateBehaviorTree(title, desc string)  *BehaviorTree {
	bt := &BehaviorTree{}
	bt.Id = CreateUUID()
	bt.Title = title
	bt.Description = desc
	return bt
}


