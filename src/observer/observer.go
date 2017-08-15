package observer

import "github.com/rookie-xy/hubble/src/prototype"

type Subject interface {
    Attach(observer Observer)
    Notify()
}

type Observer interface {
    Update(data prototype.Object) int
}

var Observers = map[string]Observer{}
var Subjects = map[string]Subject{}
