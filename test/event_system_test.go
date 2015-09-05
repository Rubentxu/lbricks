package test_lbricks


import (
	"testing"
	"github.com/Rubentxu/lbricks"
	"github.com/Rubentxu/lbricks/goflow"
	"reflect"
)


func GetPorts() [] flow.Port {
	return [] flow.Port{
		{Channel : reflect.ValueOf(make(chan *lbricks.PreloadEvent))},
		{Channel : reflect.ValueOf(make(chan *lbricks.SetupEvent))},
		{Channel : reflect.ValueOf(make(chan *lbricks.CloseEvent))},
		{Channel : reflect.ValueOf(make(chan *lbricks.UpdateEvent))},
		{Channel : reflect.ValueOf(make(chan *lbricks.RenderEvent))},
		{Channel : reflect.ValueOf(make(chan *lbricks.ResizeEvent))},
		{Channel : reflect.ValueOf(make(chan *lbricks.StepEvent))},
		{Channel : reflect.ValueOf(make(chan *lbricks.MouseEvent))},
		{Channel : reflect.ValueOf(make(chan *lbricks.ScrollEvent))},
		{Channel : reflect.ValueOf(make(chan *lbricks.KeyEvent))},
		{Channel : reflect.ValueOf(make(chan *lbricks.TypeKeyEvent))},
	}

}

func TestRegisterInputChannel(t *testing.T) {
	e_system := lbricks.CreateEventSystem(1000)
	inputs := GetPorts()

	for _,port := range inputs{
		e_system.RegisterInputChannel(port)
	}

}