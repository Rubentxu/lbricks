package test_test


import (
	"testing"
	"github.com/Rubentxu/lbricks/bgo"
	"os"
	"fmt"
)
var testMemory *bgo.Memory

func TestMain(m *testing.M) {
	// your func
	setup()

	retCode := m.Run()

	// your func
	//teardown()

	// call with result of m.Run()
	os.Exit(retCode)
}

func setup() {
	blackboard := bgo.CreateBlackboard();
	testMemory = blackboard.GetBaseMemory()

}



func BenchmarkConcurrent(b *testing.B)  {
	for i:=0; i < b.N; i++ {
		key := fmt.Sprintf("%d",i)
		go mBool(key)
	}

}

func mBool(key string) {
	testMemory.Bool[key+"A"] = true
	testMemory.Bool[key+"B"] = false
}

func TestBool(t *testing.T) {
	testMemory.Bool["A"] = true
	testMemory.Bool["B"] = false

	if testMemory.Bool["A"] != true && testMemory.Bool["B"] != false {
		t.Error("Error testBool not true")
	}

}

func TestInteger(t *testing.T) {
	testMemory.Integer["test"] = 1

	if testMemory.Integer["test"] != 1 {
		t.Error("Error testInteger")
	}

}

func TestFloat(t *testing.T) {
	testMemory.Float["test"] = 1

	if testMemory.Float["test"] != 1 {
		t.Error("Error testFloat")
	}

}

func TestComplex(t *testing.T) {
	testMemory.Complex["test"] = complex64(1)

	if testMemory.Complex["test"] != complex64(1) {
		t.Errorf("Error testComplex expected 1  --->  %c", testMemory.Complex["test"])

	}

}



