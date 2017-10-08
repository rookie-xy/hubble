package event

import "github.com/rookie-xy/hubble/types"

type Event interface {
    ID() string
    Set()
    Get() string
    Value() types.Value
}
