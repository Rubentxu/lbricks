package test_test

import (
	"testing"
	"github.com/Rubentxu/lbricks/bgo"
	"time"
)


type TestDelayNode struct {
	bgo.BaseNode
	delay time.Duration
}

func (this *TestDelayNode) Tick(context *bgo.Context) bgo.Status {
	time.Sleep(this.delay * time.Millisecond )
	return bgo.SUCCESS
}

func NewTestDelayNode(title string, delay time.Duration) *TestDelayNode {
	node := &TestDelayNode{}
	node.ID = bgo.CreateUUID()
	node.Category = bgo.ACTION
	node.Name = "TestDelayNode"
	node.Title = title
	node.delay = delay
	return node
}


func TestInverter(t *testing.T) {
	blackboard :=  bgo.CreateBlackboard()
	tree := bgo.CreateBehaviorTree("Pruebas", "Test",blackboard)
	context := bgo.CreateContext("")
	tree.Root = bgo.NewInverter("InveterTest",
									NewTestNode("Nodo1",bgo.SUCCESS),
								)

	status := tree.Tick(context)

	if status == bgo.SUCCESS {
		t.Error("Error status no es lo experado ")
	} else {
		t.Logf("Final status %s \r",status)
	}

}


func TestLimiter(t *testing.T) {
	blackboard :=  bgo.CreateBlackboard()
	tree := bgo.CreateBehaviorTree("Pruebas", "Test",blackboard)
	context := bgo.CreateContext("")
	tree.Root = bgo.NewLimiter("LimiterTest", 2,
									NewTestNode("Nodo1",bgo.SUCCESS),
								)

	status := tree.Tick(context)
	status = tree.Tick(context)
	status = tree.Tick(context)

	if status == bgo.SUCCESS {
		t.Error("Error status no es lo experado ")
	} else {
		t.Logf("Final status %s \r",status)
	}

}


func TestMaxTime(t *testing.T) {
	blackboard :=  bgo.CreateBlackboard()
	tree := bgo.CreateBehaviorTree("Pruebas", "Test",blackboard)
	context := bgo.CreateContext("")
	tree.Root = bgo.NewMaxTime("MaxTimeTest", 200,
									NewTestDelayNode("Nodo1",300),
								)

	status := tree.Tick(context)

	if status == bgo.SUCCESS {
		t.Error("Error status no es lo experado ")
	} else {
		t.Logf("Final status %s \r",status)
	}

}


func TestRepeater(t *testing.T) {
	blackboard :=  bgo.CreateBlackboard()
	tree := bgo.CreateBehaviorTree("Pruebas", "Test",blackboard)
	context := bgo.CreateContext("")
	tree.Root = bgo.NewRepeater("RepeaterTest", 3,
									NewTestNode("Nodo1",bgo.FAILURE),
								)

	tree.Tick(context)

	elem ,_ := context.GetContextMemory().ArrayString["StatusResponses"]
	var expected []string = []string {"Failure","Failure","Failure" }

	if !ArrayEquals(expected,elem) {
		t.Errorf("Error StatusResponses no son iguales a lo experado %s <---> %s", expected,elem)
	} else {
		t.Logf("Final status %s \r",elem)
	}
}


func TestNotRepeater(t *testing.T) {
	blackboard :=  bgo.CreateBlackboard()
	tree := bgo.CreateBehaviorTree("Pruebas", "Test",blackboard)
	context := bgo.CreateContext("")
	tree.Root = bgo.NewRepeater("RepeaterTest", 3,
		NewTestNode("Nodo1",bgo.RUNNING),
	)

	tree.Tick(context)

	elem ,_ := context.GetContextMemory().ArrayString["StatusResponses"]
	var expected []string = []string { "Runnig" }

	if !ArrayEquals(expected,elem) {
		t.Errorf("Error StatusResponses no son iguales a lo experado %s <---> %s", expected,elem)
	} else {
		t.Logf("Final status %s \r",elem)
	}
}


func TestRepeatUntilFailure(t *testing.T) {
	blackboard :=  bgo.CreateBlackboard()
	tree := bgo.CreateBehaviorTree("Pruebas", "Test",blackboard)
	context := bgo.CreateContext("")
	tree.Root = bgo.NewRepeatUntilFailure("RepeatUntilFailureTest", 3,
										NewTestNode("Nodo1",bgo.SUCCESS),
										)

	tree.Tick(context)

	elem ,_ := context.GetContextMemory().ArrayString["StatusResponses"]
	var expected []string = []string { "Success","Success","Success" }

	if !ArrayEquals(expected,elem) {
		t.Errorf("Error StatusResponses no son iguales a lo experado %s <---> %s", expected,elem)
	} else {
		t.Logf("Final status %s \r",elem)
	}
}


func TestNotRepeatUntilFailure(t *testing.T) {
	blackboard :=  bgo.CreateBlackboard()
	tree := bgo.CreateBehaviorTree("Pruebas", "Test",blackboard)
	context := bgo.CreateContext("")
	tree.Root = bgo.NewRepeatUntilFailure("RepeatUntilFailureTest", 3,
		NewTestNode("Nodo1",bgo.FAILURE),
	)

	tree.Tick(context)

	elem ,_ := context.GetContextMemory().ArrayString["StatusResponses"]
	var expected []string = []string { "Failure" }

	if !ArrayEquals(expected,elem) {
		t.Errorf("Error StatusResponses no son iguales a lo experado %s <---> %s", expected,elem)
	} else {
		t.Logf("Final status %s \r",elem)
	}
}



func TestRepeatUntilSuccess(t *testing.T) {
	blackboard :=  bgo.CreateBlackboard()
	tree := bgo.CreateBehaviorTree("Pruebas", "Test",blackboard)
	context := bgo.CreateContext("")
	tree.Root = bgo.NewRepeatUntilSuccess("RepeatUntilSuccessTest", 3,
		NewTestNode("Nodo1",bgo.FAILURE),
	)

	tree.Tick(context)

	elem ,_ := context.GetContextMemory().ArrayString["StatusResponses"]
	var expected []string = []string { "Failure","Failure","Failure" }

	if !ArrayEquals(expected,elem) {
		t.Errorf("Error StatusResponses no son iguales a lo experado %s <---> %s", expected,elem)
	} else {
		t.Logf("Final status %s \r",elem)
	}
}


func TestNotRepeatUntilSuccess(t *testing.T) {
	blackboard :=  bgo.CreateBlackboard()
	tree := bgo.CreateBehaviorTree("Pruebas", "Test",blackboard)
	context := bgo.CreateContext("")
	tree.Root = bgo.NewRepeatUntilSuccess("RepeatUntilSuccessTest", 3,
		NewTestNode("Nodo1",bgo.SUCCESS),
	)

	tree.Tick(context)

	elem ,_ := context.GetContextMemory().ArrayString["StatusResponses"]
	var expected []string = []string { "Success" }

	if !ArrayEquals(expected,elem) {
		t.Errorf("Error StatusResponses no son iguales a lo experado %s <---> %s", expected,elem)
	} else {
		t.Logf("Final status %s \r",elem)
	}
}