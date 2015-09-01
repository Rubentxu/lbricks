package testlb

import (
	"testing"
	"github.com/Rubentxu/lbricks"
	"github.com/trustmaster/goflow"
)

func TestMouseSensor(t *testing.T) {
	d := lbricks.NewMouseSensor("test",1,lbricks.LEFT_BUTTON_UP)
	in := make(chan *lbricks.MouseEvent)
	out := make(chan interface{})
	d.MouseEvent = in
	d.Out = out
	flow.RunProc(d)
	for i := 0.0; i < 10000000.0; i++ {
		var ms = &lbricks.MouseEvent{float32(i * 2.0), float32(i * 2.0), lbricks.LEFT_BUTTON_UP}
		in <- ms
		i2 := (<-out).(*lbricks.MouseEvent)
		ix2 := i * 2.0
		if i2.PosX != float32(ix2) {
			t.Errorf("Error en el test %f != %f", i2, ix2)
		}

	}
	// Shutdown the component
	close(in)
}
