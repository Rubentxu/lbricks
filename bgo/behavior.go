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
	*Blackboard

}

func (this *BehaviorTree) Tick(context *Context) Status {
	context.BehaviorTree = this
	state := ExecuteNode(this.Root,context)

	var lastOpenNodes []Node
	if nodes, ok :=this.GetTreeMemory().ArrayNode["openNodes"];ok {
		lastOpenNodes = nodes
	}

	var openNodes []Node


	for _,currNode := range context.openNodes {
		openNodes = append(openNodes, currNode)
		for _,lastNode := range lastOpenNodes {
			if lastNode.Id() != currNode.Id() {
				lastNode.Close(context)
			}
		}
	}
	this.GetTreeMemory().ArrayNode["openNodes"] = openNodes
	this.GetTreeMemory().Integer["nodeCount"] = context.nodeCount
	return state

}

func (this *BehaviorTree) GetTreeMemory() *Memory {
	return this.getTreeMemory(this.Id)
}

func  (this *BehaviorTree) GetNodeMemory(node Node) *Memory {
	return this.getNodeMemory(this.Id, node.Id())
}

func CreateBehaviorTree(title, desc string,blackboard *Blackboard)  *BehaviorTree {
	bt := &BehaviorTree{}
	bt.Id = CreateUUID()
	bt.Title = title
	bt.Description = desc
	bt.Blackboard = blackboard
	return bt
}


