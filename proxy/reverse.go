package proxy

import (
    "github.com/rookie-xy/hubble/src/event"
    "github.com/rookie-xy/hubble/src/log"
    "github.com/rookie-xy/hubble/src/types"
)

type Server func(log.Log, types.Value) (Reverse, error)

type Reverse interface {
    Insert(e adapter.Event) int
    Update(e adapter.Event) int
    Delete(e adapter.Event) int
    Select(e adapter.Event) int
    Close() int
}

var Reverses = map[string]Server{}