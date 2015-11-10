package test_test

import (
	"testing"
	"github.com/Rubentxu/lbricks/bgo"

)

type TestNode struct {
	bgo.BaseNode
	statusResponse  bgo.Status
}

func (this *TestNode) Tick(context *bgo.Context) bgo.Status {
	var statusResponses []bgo.Status
	elem ,ok := context.Blackboard.Get("StatusResponses",context.Tree.Id, "")
	if !ok {
		statusResponses = make([]bgo.Status,0,1)
	} else {
		statusResponses = elem.([]bgo.Status)
	}
	statusResponses = append(statusResponses,this.statusResponse)
	context.Blackboard.Set("StatusResponses",statusResponses,context.Tree.Id, "")
	return this.statusResponse
}

func NewTestNode(title string, statusResponse bgo.Status) *TestNode {
	node := &TestNode{}
	node.ID = bgo.CreateUUID()
	node.Category = bgo.ACTION
	node.Name = "TestNode"
	node.Title = title
	node.statusResponse = statusResponse
	return node
}


func ArrayEquals(a []bgo.Status, b []bgo.Status) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestSequenceRoot(t *testing.T) {
	tree := bgo.CreateBehaviorTree("Pruebas", "Test")
	context := bgo.CreateContext("")

	tree.Root  = bgo.NewSequence("SequenceTest",
								bgo.NewPriority("PriorityTest",
												NewTestNode("Nodo1",bgo.FAILURE),
												NewTestNode("Nodo2",bgo.SUCCESS),
												NewTestNode("Nodo3",bgo.SUCCESS)),

								NewTestNode("Nodo4",bgo.SUCCESS),
								NewTestNode("Nodo5",bgo.SUCCESS))


	tree.Tick(context)
	elem ,_ := context.Blackboard.Get("StatusResponses",context.Tree.Id, "")
	var expected []bgo.Status = []bgo.Status { bgo.FAILURE,bgo.SUCCESS,bgo.SUCCESS,bgo.SUCCESS }

	if !ArrayEquals(expected,elem.([]bgo.Status)) {
		t.Error("Error StatusResponses no son iguales a lo experado ", expected)
	} else {
		t.Logf("Final status %s \r",elem.([]bgo.Status))
	}

}


func TestPriorityRoot(t *testing.T) {
	tree := bgo.CreateBehaviorTree("Pruebas2", "Test2")
	context := bgo.CreateContext("")

	tree.Root  =	bgo.NewPriority("PriorityTest",
									bgo.NewSequence("SequenceTestA",
												NewTestNode("NodoA",bgo.SUCCESS),
												NewTestNode("NodoB",bgo.FAILURE),
												NewTestNode("NodoC",bgo.SUCCESS)),
									bgo.NewSequence("SequenceTestB",
													NewTestNode("NodoD",bgo.SUCCESS),
													NewTestNode("NodoE",bgo.SUCCESS)),
									)

	tree.Tick(context)
	elem ,_ := context.Blackboard.Get("StatusResponses",context.Tree.Id, "")
	var expected []bgo.Status = []bgo.Status { bgo.SUCCESS,bgo.FAILURE,bgo.SUCCESS,bgo.SUCCESS }

	if  !ArrayEquals(expected,elem.([]bgo.Status)) {
		t.Error("Error StatusResponses no son iguales a lo experado",expected)
	} else {
		t.Logf("Final status %s \r ",elem.([]bgo.Status))
	}


}
