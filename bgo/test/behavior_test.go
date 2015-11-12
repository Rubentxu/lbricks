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
	var statusResponses []string = context.GetTreeMemory().ArrayString["StatusResponses"]
	statusResponses = append(statusResponses,string(this.statusResponse))
	context.GetTreeMemory().ArrayString["StatusResponses"] = statusResponses
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


func ArrayEquals(a []string, b []string) bool {
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
	blackboard :=  bgo.CreateBlackboard()
	tree := bgo.CreateBehaviorTree("Pruebas", "Test", blackboard)
	context := bgo.CreateContext("")

	tree.Root  = bgo.NewSequence("SequenceTest",
								bgo.NewPriority("PriorityTest",
												NewTestNode("Nodo1",bgo.FAILURE),
												NewTestNode("Nodo2",bgo.SUCCESS),
												NewTestNode("Nodo3",bgo.SUCCESS)),

								NewTestNode("Nodo4",bgo.SUCCESS),
								NewTestNode("Nodo5",bgo.SUCCESS))


	tree.Tick(context)
	elem ,_ := context.GetTreeMemory().ArrayString["StatusResponses"]
	expected := []string { "Failure","Succes","Succes","Succes" }

	if !ArrayEquals(expected,elem) {
		t.Errorf("Error StatusResponses no son iguales a lo experado %s <---> %s", expected,elem)
	} else {
		t.Logf("Final status %s \r",elem)
	}

}


func TestPriorityRoot(t *testing.T) {
	blackboard :=  bgo.CreateBlackboard()
	tree := bgo.CreateBehaviorTree("Pruebas2", "Test2", blackboard)
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

	elem ,_ := context.GetTreeMemory().ArrayString["StatusResponses"]
	expected := []string { "Succes","Failure","Succes","Succes" }

	if !ArrayEquals(expected,elem) {
		t.Errorf("Error StatusResponses no son iguales a lo experado %s <---> %s", expected,elem)
	} else {
		t.Logf("Final status %s \r",elem)
	}


}
