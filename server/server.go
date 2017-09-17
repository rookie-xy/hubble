package server

import (
    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/hubble/log"
)

type Factory func(log.Log, types.Value) (Server, error)

type Server interface {
    Listen()
    Accept()
}

var Servers = map[string]Factory{}
