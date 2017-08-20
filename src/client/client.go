package client

import (
    "github.com/rookie-xy/hubble/src/event"
    "github.com/rookie-xy/hubble/src/log"
    "github.com/rookie-xy/hubble/src/types"
)

type Factory func(log.Log, types.Value) (Client, error)

type Client interface {
    Sender(e event.Event) int
    Close() int
}

var Clients = map[string]Factory{}
