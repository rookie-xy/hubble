package observer

import "github.com/rookie-xy/worker/src/prototype"

type Subject interface {
    Attach(observer Observer)
    Notify()
}

type Observer interface {
    Update(data prototype.Object) int
}

var Observers = map[string]Observer{}
var Subjects = map[string]Subject{}

/*
type Topic struct {
   	observers []Observer
   	Data chan prototype.Object
}

func (r *Topic) Attach(observer Observer) {
   	r.observers = append(r.observers, observer)
}

func (r *Topic) Notify() {
   	for _, observer := range r.observers {
		      observer.Update(r.Data)
	   }
}


type ConcreteObserver struct {
   	state string
}

func (r *ConcreteObserver) Update(state string) {
   	r.state = state
}
*/
