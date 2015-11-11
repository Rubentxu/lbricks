package test_test


import (
	"testing"
	"github.com/Rubentxu/lbricks/bgo"

)


func TestGetInt8(t *testing.T) {
	blackboard := bgo.CreateBlackboard();
	baseMemory := blackboard.GetBaseMemory()

	baseMemory.Byte["testBool"] = true

	if baseMemory.Byte["testBool"] != true {
		t.Error("Error testBool not true")
	}
}