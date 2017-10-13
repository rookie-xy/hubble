package output

import (
    "github.com/rookie-xy/hubble/proxy"
    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/types"
)

type Factory func(log.Log, types.Value) (Output, error)

type Output interface {
    proxy.Forward
}

var Outputs = map[string]Factory{}
