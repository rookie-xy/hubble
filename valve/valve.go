package valve

import (
    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/hubble/log"
//    "github.com/rookie-xy/hubble/event"
)

type Factory func(log.Log, types.Value) (Valve, error)

// chain_of_responsibility
type Valve interface {
    Filter(text string) bool
    Next(Valve)
}

var Valves = map[string]Factory{}
