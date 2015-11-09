package test_test

import (
	"testing"
	"github.com/Rubentxu/lbricks/bgo"
	"fmt"

)

type TestNode struct {
	bgo.BaseNode

}

func (this *TestNode) Tick(context *bgo.Context) bgo.Status {
	fmt.Println("Test Succes %s", this.Title)
	return bgo.SUCCESS
}

func NewTestNode(title string) *TestNode {
	wait := &TestNode{}
	wait.ID = bgo.CreateUUID()
	wait.Category = bgo.ACTION
	wait.Name = "TestNode"
	wait.Title = title
	return wait
}


func TestSignal(t *testing.T) {
	tree := bgo.CreateBehaviorTree("Pruebas", "Test")
	blackboard := bgo.CreateBlackboard();



	tree.Root  = bgo.NewSequence("SequenceTest",
								bgo.NewPriority("PriorityTest",
												bgo.CreateBaseNode("base"),
												NewTestNode("Nodo1"),
												NewTestNode("Nodo2")),

								NewTestNode("Nodo3"),
								NewTestNode("Nodo4"))


	status := tree.Tick("",blackboard)
	t.Logf("Final status %s ",status)
}
