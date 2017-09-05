package observer

import (
    "github.com/rookie-xy/hubble/types"
)

type Subject interface {
    Attach(observer Observer)
    Notify(o types.Object)
}

var Subjects = map[string]Subject{}
