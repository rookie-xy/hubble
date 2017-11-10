package proxy

import (
    "github.com/rookie-xy/hubble/event"
    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/types"
)

type Client func(log.Log, types.Value) (Forward, error)

type Forward interface {
    Sender(event.Event) error
    Close()
}

var Clients = map[string]Client{}
var Forwards = map[string]Forward{}
