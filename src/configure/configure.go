package configure

import (
    "github.com/rookie-xy/hubble/src/types"
)

const Name = "configure"

var Event chan []byte = make(chan []byte)

type Configure struct {
    Iterms  []*types.Iterm
}

type ConfigureIterator struct {
   *Configure
    index    int
    internal int
}

func New() *Configure {
    return &Configure{}
}

func (r *Configure) Iterator() *ConfigureIterator {
    return &ConfigureIterator{Configure: r}
}

func (r *Configure) Add(key, value types.Value) {
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
