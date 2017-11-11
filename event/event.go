package event

import (
    "github.com/rookie-xy/hubble/state"
)

type Event interface {
    state.State
}
