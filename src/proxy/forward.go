package proxy

import (
    "github.com/rookie-xy/hubble/src/event"
    "github.com/rookie-xy/hubble/src/log"
    "github.com/rookie-xy/hubble/src/types"
)

type Client func(log.Log, types.Value) (Forward, error)

type Forward interface {
    Sender(e event.Event) int
    Close() int
}

var Forwards = map[string]Client{}
