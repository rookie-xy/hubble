package observer

import (
    "github.com/rookie-xy/hubble/src/types"
)

type Subject interface {
    Attach(observer Observer)
    Notify()
}

type Observer interface {
    Update(v types.Value) int
}

var Observers = map[string]Observer{}
var Subjects = map[string]Subject{}
