package configure

import (
    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/observer"
)

type Configure struct {
    log.Log
    observers  []observer.Observer
    Iterms     []*types.Iterm
}


