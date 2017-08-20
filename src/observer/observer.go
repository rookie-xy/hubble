package observer

import (
    "github.com/rookie-xy/hubble/src/types"
)

type Subject interface {
    Attach(observer Observer)
    Notify(o types.Object)
}

type Observer interface {
    Update(o types.Object) int
}

var Observers = map[string]Observer{}
var Subjects = map[string]Subject{}
