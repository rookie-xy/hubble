package client

import (
    "github.com/rookie-xy/hubble/src/event"
    "github.com/rookie-xy/hubble/src/log"
    "github.com/rookie-xy/hubble/src/command"
)

type Factory func(log.Log, *command.Command) (Client, error)

type Client interface {
    Sender(e event.Event) int
    Close() int
}

var Clients = map[string]Factory{}
