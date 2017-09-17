package client

import (
    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/hubble/log"
)

type Factory func(log.Log, types.Value) (Client, error)

type Client interface {
    Connect()
}

var Clients = map[string]Factory{}
