package proxy

import (
    "github.com/rookie-xy/hubble/src/event"
    "github.com/rookie-xy/hubble/src/log"
    "github.com/rookie-xy/hubble/src/types"
)

type Server func(log.Log, types.Value) (Reverse, error)

type Reverse interface {
    Insert(e event.Event) int
    Update(e event.Event) int
    Delete(e event.Event) int
    Select(e event.Event) int
    Close() int
}

var Reverses = map[string]Server{}