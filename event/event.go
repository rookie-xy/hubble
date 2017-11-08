package event

import (
    "github.com/rookie-xy/hubble/state"
)

type Event interface {
//    file.Status
    state.State
}

/*
type Event interface {
    ID() string
    GetHeader() types.Map
    GetBody() Message
    GetState() file.State
}

type Message interface {
    ID() uint64
    GetContent() []byte
    Json() string
}
*/
