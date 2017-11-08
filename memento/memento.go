package memento

var Name = ""
/*
import "testing"

type Originator struct {
   	State string
}

func (r *Originator) CreateMemento() *Memento {
   	return &Memento{file: r.State}
}

func (r *Originator) SetMemento(memento *Memento) {
   	r.State = memento.GetState()
}

type Memento struct {
   	file string
}

func (r *Memento) GetState() string {
   	return r.file
}

type Caretaker struct {
   	Memento *Memento
}

func TestMomento(t *testing.T) {
   	originator := &Originator{}
   	caretaker := &Caretaker{}

   	originator.State = "On"

   	caretaker.Memento = originator.CreateMemento()

   	originator.State = "Off"

	   originator.SetMemento(caretaker.Memento)

   	if originator.State != "On" {
  		    t.Errorf("Expect State to %s, but %s", originator.State, "On")
	   }
}
*/
