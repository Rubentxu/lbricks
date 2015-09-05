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
	eventChan := make(chan *lbricks.PreloadEvent)
	inputs := flow.Port{Channel : reflect.ValueOf(eventChan)}

	e_system.RegisterInputChannel(inputs)
	go e_system.Preload()

	event:= <-eventChan
	t.Logf("Event recibido: ",event)
}

func TestRegisterInputChannels(t *testing.T) {
	e_system := lbricks.CreateEventSystem(1000)
	inputs := GetPorts()
	for _,port := range inputs {
		e_system.RegisterInputChannel(port)
	}

	go e_system.Preload()
	go e_system.Setup()
	go e_system.Close()

	<-inputs[0].Channel.Interface().(chan *lbricks.PreloadEvent)
	<-inputs[1].Channel.Interface().(chan *lbricks.SetupEvent)
	<-inputs[2].Channel.Interface().(chan *lbricks.CloseEvent)

	t.Logf("Event recibidoS: ")
}
