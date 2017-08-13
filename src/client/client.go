package client

import "github.com/rookie-xy/worker/src/event"

type Factory func(name string) (Client, error)

type Client interface {
    Sender(e event.Event) int
    Close()
}

type Configure struct {
    name     string
    content  string
}

var Clients = map[string]Factory{}
