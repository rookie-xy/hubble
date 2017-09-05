package command

import (
    "github.com/rookie-xy/hubble/types"
)

type SetFunc func(cmd *Item, meta *Command, val types.Object) int

type Item struct {
    Command  *Command
    Type      int
    Scope     string
    Set       SetFunc
    State     bool
    Offset    uintptr
    Load      types.Object
}

var Pool []Item
