package test_test

import (
	"testing"
	"github.com/Rubentxu/lbricks/bgo"

)


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