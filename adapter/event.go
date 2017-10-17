package adapter

import "github.com/rookie-xy/hubble/event"

type MessageEvent interface {
    event.Event

    ID() uint64
    GetContent() []byte
    Json() string
}

func ToMessageEvent(e event.Event) MessageEvent {
    return e.(MessageEvent)
}
