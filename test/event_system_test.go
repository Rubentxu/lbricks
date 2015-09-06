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

	e_system.RegisterEntityInput(inputs)
	go e_system.Preload()

	event:= <-eventChan
	t.Logf("Event recibido: ",event)
}

func TestRegisterInputChannels(t *testing.T) {
	e_system := lbricks.CreateEventSystem(1000)
	inputs := GetPorts()
	for _,port := range inputs {
		e_system.RegisterEntityInput(port)
	}

	go e_system.Preload()
	go e_system.Setup()
	go e_system.Close()
	go e_system.Update(0.5)
	go e_system.Render()
	go e_system.Resize(100,100)
	go e_system.Step(100.0, 100)
	go e_system.Mouse(10,10,3)
	go e_system.Scroll(10)
	go e_system.Key(65,0,0x0001)

	<-inputs[0].Channel.Interface().(chan *lbricks.PreloadEvent)
	<-inputs[1].Channel.Interface().(chan *lbricks.SetupEvent)
	<-inputs[2].Channel.Interface().(chan *lbricks.CloseEvent)
	update := <-inputs[3].Channel.Interface().(chan *lbricks.UpdateEvent)
	if update.DeltaTime != 0.5 {
		t.Errorf("Error update deltatime")
	}
	<-inputs[4].Channel.Interface().(chan *lbricks.RenderEvent)
	resize := <-inputs[5].Channel.Interface().(chan *lbricks.ResizeEvent)
	if resize.Height != 100 || resize.Width != 100 {
		t.Errorf("Error resize event")
	}
	stepEvent := <-inputs[6].Channel.Interface().(chan *lbricks.StepEvent)
	if stepEvent.Step != 100.0 || stepEvent.NumStep != 100 {
		t.Errorf("Error step event")
	}
	mouseEvent := <-inputs[7].Channel.Interface().(chan *lbricks.MouseEvent)
	if mouseEvent.PosX != 10 || mouseEvent.PosY != 10 ||
		mouseEvent.Action != lbricks.RIGHT_BUTTON_DOWN{
		t.Errorf("Error mouse event", mouseEvent.Action)
	}
	scrollEvent := <-inputs[8].Channel.Interface().(chan *lbricks.ScrollEvent)
	if scrollEvent.AmountScroll < 10 || scrollEvent.Action != lbricks.WHEEL_UP{
		t.Errorf("Error scroll event", mouseEvent.Action)
	}
	/*keyEvent := <-inputs[9].Channel.Interface().(chan *lbricks.KeyEvent)
	if keyEvent.Key != lbricks.A || keyEvent.Action != lbricks.PRESS ||
		 keyEvent.Modifier != lbricks.SHIFT{
		t.Errorf("Error key event", keyEvent.Modifier)
	}*/
}
