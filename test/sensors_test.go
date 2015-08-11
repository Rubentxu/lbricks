package testlb

import (
	"testing"

	"github.com/Rubentxu/lbricks"
	"github.com/ajhager/engi"
	"github.com/trustmaster/goflow"
)

func TestMouseInput(t *testing.T) {
	d := new(lbricks.MouseSensor)
	in := make(chan *lbricks.MouseSignal)
	out := make(chan *lbricks.MouseSignal)
	d.In = in
	d.Out = out
	flow.RunProc(d)
	for i := 0.0; i < 1000000.0; i++ {
		var ms = &lbricks.MouseSignal{float32(i * 2.0), float32(i * 2.0), engi.PRESS}
		in <- ms
		i2 := <-out
		ix2 := i * 2
		if i2.PosX != float32(ix2) {
			t.Errorf("Error en el test %f != %f", i2, ix2)
		}

	}
	// Shutdown the component
	close(in)
}
