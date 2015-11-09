package test_test

import (
	"testing"
	"github.com/Rubentxu/lbricks/bgo"
	"fmt"

)

type TestNode struct {
	bgo.BaseNode

}

func (this *TestNode) tick(tick bgo.Tick) bgo.Status {
	fmt.Println("Test Succes %s", this.Id)
	return bgo.SUCCESS
}

func NewTestNode() *TestNode {
	wait := &TestNode{}
	wait.SetId(bgo.CreateUUID())
	wait.Category = bgo.ACTION
	wait.Name = "TestNode"
	return wait
}


func TestSignal(t *testing.T) {
	tree := bgo.CreateBehaviorTree("Pruebas", "Test")
	blackboard := bgo.CreateBlackboard();



	tree.Root  = bgo.NewPriority("PriorityTest",
								bgo.NewSequence("SequenceTest",
												NewTestNode(),
												NewTestNode(),
												NewTestNode()),

				)

	status := tree.Tick("",blackboard)
	t.Logf("Final status %s",status)
}
