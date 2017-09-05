package observer

import (
    "github.com/rookie-xy/hubble/types"
)

type Observer interface {
    Update(o types.Object) int
}

var Observers = map[string]Observer{}
