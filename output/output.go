package output

import (
    "github.com/rookie-xy/hubble/proxy"
    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/types"
)

type Factory func(log.Log, types.Value) (Output, error)

type Output interface {
    //proxy.Forward
    Connect() proxy.Forward
    Accept() pipeline.Queue
}

var Outputs = map[string]Factory{}
