package configure

import (
    "fmt"

    "github.com/rookie-xy/hubble/src/types"
    "github.com/rookie-xy/hubble/src/module"
    "github.com/rookie-xy/hubble/src/log"
    "github.com/rookie-xy/hubble/src/observer"
    "github.com/rookie-xy/hubble/src/state"
)

type Configure struct {
    log.Log
    observers  []observer.Observer
    Iterms     []*types.Iterm
}

type ConfigureIterator struct {
   *Configure

    index    int
    internal int
}

type build func(name string, i types.Iterator, l module.Load) int

var Build build = nil

func New(log log.Log) *Configure {
    return &Configure{
        Log:log,
    }
}

func (r *Configure) Attach(o observer.Observer) {
    if o != nil {
        r.observers = append(r.observers, o)
        return
    }

    fmt.Println("attach error")
    return
}

func (r *Configure) Notify(o types.Object) {
    for _, observer := range r.observers {
        if observer.Update(o) == state.Error {
            break
        }
    }
}

func (r *Configure) Iterator() *ConfigureIterator {
    return &ConfigureIterator{Configure: r}
}

func (r *Configure) Add(key types.Value, value types.Object) {
    iterm := &types.Iterm{key, value}
    r.Iterms = append(r.Iterms, iterm)
}

func (r *ConfigureIterator) Index() int {
    return r.index
}

func (r *ConfigureIterator) Iterm() *types.Iterm {
    if iterm := r.Iterms[r.index]; iterm != nil {
        return iterm
    }

    return nil
}

func (r *ConfigureIterator) Has() bool {
    if r.internal < 0 || r.internal >= len(r.Iterms) {
        return false
    }

    return true
}

func (r *ConfigureIterator) Next() {
    r.internal++
    if r.Has() {
        r.index++
    }
}

func (r *ConfigureIterator) Prev() {
    r.internal--
    if r.Has() {
        r.index--
    }
}

func (r *ConfigureIterator) Reset() {
    r.index = 0
    r.internal = 0
}

func (r *ConfigureIterator) End() {
    r.index = len(r.Iterms) - 1
    r.internal = r.index
}
