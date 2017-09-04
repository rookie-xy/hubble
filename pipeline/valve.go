package pipeline

import "github.com/rookie-xy/hubble/src/event"

// chain_of_responsibility
type Valve interface {
    Filter(event.Event) bool
    Next(Valve)
}
