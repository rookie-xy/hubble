package adapter

import (
    "github.com/rookie-xy/hubble/proxy"
    "github.com/rookie-xy/hubble/event"
)

type BatchForward interface {
    proxy.Forward
    Commit(event.Event) bool
    Senders() ([]event.Event, error)
}

func ToBatchForward(f proxy.Forward) BatchForward {
    return f.(BatchForward)
}
