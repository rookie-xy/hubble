package input

import (
    "io"
    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/hubble/source"
)

type Factory func(log.Log, types.Value) (Input, error)

type Input interface {
    Init(source source.Source) error
    io.ReadCloser
}

var Inputs = map[string]Factory{}
