package test_test


import (
	"testing"
	"github.com/Rubentxu/lbricks/bgo"

)


func TestGetInt8(t *testing.T) {
	blackboard := bgo.CreateBlackboard();
	baseMemory := blackboard.GetBaseMemory()

	baseMemory.Bool["testBool"] = true
	baseMemory.Bool["testBool2"] = false

	if baseMemory.Bool["testBool"] != true && baseMemory.Bool["testBool2"] != false {
		t.Error("Error testBool not true")
	}

}