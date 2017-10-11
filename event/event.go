package event

import "github.com/rookie-xy/hubble/types"

type Event interface {
    ID() string
    GetHeader() types.SiMap
    GetBody() Message
    GetFooter() []byte
}

type Message interface {
    ID() uint64
    GetContent() []byte
    Json() string
}
