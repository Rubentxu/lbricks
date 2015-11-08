package tests_test

import (
	"testing"
	"github.com/Rubentxu/lbricks"
	"fmt"
	"time"
)


func TestSignal(t *testing.T) {
	lbricks.FromValues(1, 2, 3, 4, 5, 6, 7,8,9,10,11,12).
	Filter(func(value interface{}) bool {
		return value.(int) % 3 == 0
	}).
	Map(func(value interface{}) interface{} {
		return value.(int) * 2
	}).
	Subscribe(func (value interface {}) {
		fmt.Println("From Subscribe ",value.(int))
	})

	time.Sleep(time.Nanosecond * 100000)
}

