package event

import (
    "github.com/rookie-xy/hubble/state"
)

type Event interface {
    state.Status
}

/*
type Event interface {
    ID() string
    GetHeader() types.Map
    GetBody() Message
    GetState() state.State
}

type Message interface {
    ID() uint64
    GetContent() []byte
    Json() string
}
*/
