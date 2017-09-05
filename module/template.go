package module

import (
    "github.com/rookie-xy/hubble/log"
)

// template
type Factory func(log log.Log) Template

type Template interface {
    Init()
    Main()
    Exit(code int)
}

var Pool = map[string]*Factory{}
