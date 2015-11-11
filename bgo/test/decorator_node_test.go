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
	tree := bgo.CreateBehaviorTree("Pruebas", "Test")
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
	tree := bgo.CreateBehaviorTree("Pruebas", "Test")
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
	tree := bgo.CreateBehaviorTree("Pruebas", "Test")
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
	tree := bgo.CreateBehaviorTree("Pruebas", "Test")
	context := bgo.CreateContext("")
	tree.Root = bgo.NewRepeater("RepeaterTest", 3,
									NewTestNode("Nodo1",bgo.FAILURE),
								)

	status := tree.Tick(context)

	elem ,_ := context.Blackboard.Get("StatusResponses",context.Tree.Id, "")
	var expected []bgo.Status = []bgo.Status { bgo.FAILURE,bgo.FAILURE,bgo.FAILURE }

	if  !ArrayEquals(expected,elem.([]bgo.Status)) {
		t.Errorf("Error StatusResponses no son iguales a lo experado %s, %s",expected,status)
	} else {
		t.Logf("Final status %s \r ",elem.([]bgo.Status))
	}
}


func TestNotRepeater(t *testing.T) {
	tree := bgo.CreateBehaviorTree("Pruebas", "Test")
	context := bgo.CreateContext("")
	tree.Root = bgo.NewRepeater("RepeaterTest", 3,
		NewTestNode("Nodo1",bgo.RUNNING),
	)

	status := tree.Tick(context)

	elem ,_ := context.Blackboard.Get("StatusResponses",context.Tree.Id, "")
	var expected =  []bgo.Status { bgo.RUNNING }

	if  !ArrayEquals(expected,elem.([]bgo.Status)) {
		t.Errorf("Error StatusResponses no son iguales a lo experado %s, %s",expected,status)
	} else {
		t.Logf("Final status %s \r ",elem.([]bgo.Status))
	}
}


func TestRepeatUntilFailure(t *testing.T) {
	tree := bgo.CreateBehaviorTree("Pruebas", "Test")
	context := bgo.CreateContext("")
	tree.Root = bgo.NewRepeatUntilFailure("RepeatUntilFailureTest", 3,
										NewTestNode("Nodo1",bgo.SUCCESS),
										)

	status := tree.Tick(context)

	elem ,_ := context.Blackboard.Get("StatusResponses",context.Tree.Id, "")
	var expected []bgo.Status = []bgo.Status { bgo.SUCCESS,bgo.SUCCESS,bgo.SUCCESS }

	if  !ArrayEquals(expected,elem.([]bgo.Status)) {
		t.Errorf("Error StatusResponses no son iguales a lo experado %s, %s",expected,status)
	} else {
		t.Logf("Final status %s \r ",elem.([]bgo.Status))
	}
}


func TestNotRepeatUntilFailure(t *testing.T) {
	tree := bgo.CreateBehaviorTree("Pruebas", "Test")
	context := bgo.CreateContext("")
	tree.Root = bgo.NewRepeatUntilFailure("RepeatUntilFailureTest", 3,
		NewTestNode("Nodo1",bgo.FAILURE),
	)

	status := tree.Tick(context)

	elem ,_ := context.Blackboard.Get("StatusResponses",context.Tree.Id, "")
	var expected []bgo.Status = []bgo.Status { bgo.FAILURE }

	if  !ArrayEquals(expected,elem.([]bgo.Status)) {
		t.Errorf("Error StatusResponses no son iguales a lo experado %s, %s",expected,status)
	} else {
		t.Logf("Final status %s \r ",elem.([]bgo.Status))
	}
}