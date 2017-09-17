package proxy

import (
    "github.com/rookie-xy/hubble/event"
    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/types"
)

type Server func(log.Log, types.Value) (Reverse, error)

type Reverse interface {
    Post(e event.Event) int
    Delete(e event.Event) int
    Put(e event.Event) int
    Get(e event.Event) types.Object
}

var Reverses = map[string]Server{}