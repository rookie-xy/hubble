package filter

import (
    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/hubble/log"
//    "github.com/rookie-xy/hubble/event"
)

type Factory func(log.Log, types.Value) (Filter, error)

// chain_of_responsibility
type Filter interface {
    Handler(text string) bool
    Next(Filter)
}

var Filters = map[string]Factory{}
